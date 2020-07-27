
func process(r io.Reader) error {
	_, err := ioutil.ReadAll(r)
	if err != nil {
		// DO NOT
		// return errors.New("read failed: " + err.Error())

		// instead, wrap errors:
		return errors.Wrap(err, "read failed")
	}

	// ... rest of the processing
}

func main() {
	r := // obtain the io.Reader here
	err := process(r)
	if err != nil {
		// let's check the original cause of the error
		causeErr := errors.Cause(err)
		
		// do something with causeErr
	}
}

    
