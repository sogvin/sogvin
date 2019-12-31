package errhandling

/*
func CopyFile(src, dst string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("copy %s %s: %v", src, dst, err)
		}
	}()

	r := try(os.Open(src))
	defer r.Close()

	w := try(os.Create(dst))
	defer func() {
		w.Close()
		if err != nil {
			os.Remove(dst) // only if a “try” fails
		}
	}()

	try(io.Copy(w, r))
	try(w.Close())
	return nil
}
*/
