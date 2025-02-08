package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// SecurityMiddleware applies security headers to HTTP responses
func SecurityMiddleware() gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny:             true,                  // Prevent clickjacking
		ContentTypeNosniff:    true,                  // Prevent MIME-sniffing
		BrowserXssFilter:      true,                  // Enable XSS protection in browsers
		ContentSecurityPolicy: "default-src 'self';", // Restrict content sources
		ReferrerPolicy:        "no-referrer",         // Control referrer information
		FeaturePolicy:         "vibrate 'none';",     // Limit browser features
		// SSLRedirect:           true,
		STSSeconds:           31536000, // HSTS max-age in seconds
		STSIncludeSubdomains: true,     // Apply HSTS to subdomains
		STSPreload:           true,     // Allow HSTS preload
	})

	return func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatus(400) // Bad request if security headers can't be applied
			return
		}

		c.Next()
	}
}
