package factorymethod

func ExampleFactoryMethod() {
	doLog(&FileLoggerFactory{})
	doLog(&ConsoleLoggerFactory{})
	// Output:
	// Log to file: This is a test log.
	// Log to console: This is a test log.
}

func doLog(factory LoggerFactory) {
	factory.CreateLogger().Log("This is a test log.")
}
