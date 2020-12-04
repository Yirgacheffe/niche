package main

// UseLog demonstrates using our global log
func UseLog() error {
	if err := Init(); err != nil {
		return err
	}

	WithField("key", "value").Debug("hello")
	Debug("test")

	return nil // nothing happened as normal
}
