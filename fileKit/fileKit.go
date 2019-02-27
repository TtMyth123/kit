package fileKit

func CreateFile(c []byte, fileName string) error {
	file, err := os.OpenFile(
		fileName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(c)
	return err
}
