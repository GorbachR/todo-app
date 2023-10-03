package middleware

//const TIMEOUT = 5 * time.Second
//
//func Context() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//		finished := make(chan struct{})
//		panicChan := make(chan interface{}, 1)
//
//		ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
//		defer cancel()
//		c.Request = c.Request.WithContext(ctx)
//
//		go func() {
//			defer func() {
//				if p := recover(); p != nil {
//					panicChan <- p
//				}
//			}()
//			c.Next()
//			finished <- struct{}{}
//		}()
//
//		select {
//		case <-panicChan:
//		case <-finished:
//
//		case <-ctx.Done():
//		}
//	}
//}
