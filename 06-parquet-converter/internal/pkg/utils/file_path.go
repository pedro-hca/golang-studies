package utils

func GetCsvFilePath() string {
	// absPath, err := filepath.Abs("../../.env")
	// if err != nil {
	// 	log.Fatalf("Error finding absolute path: %v", err)
	// }
	// fmt.Println("Absolute path:", absPath)

	// err = godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	// return os.Getenv("COMMON_PATH") + "golang-studies/06-parquet-converter/internal/pkg/data/csv/"
	return "/home/pedro/dev/golang-studies/06-parquet-converter/internal/pkg/data/csv/"
}
func GetJsonFilePath() string {
	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	// return os.Getenv("COMMON_PATH") + "golang-studies/06-parquet-converter/internal/pkg/data/json/"
	return "/home/pedro/dev/golang-studies/06-parquet-converter/internal/pkg/data/json/"
}
func GetParquetFilePath() string {
	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	// return os.Getenv("COMMON_PATH") + "golang-studies/06-parquet-converter/internal/pkg/data/parquet/"
	return "/home/pedro/dev/golang-studies/06-parquet-converter/internal/pkg/data/parquet/"
}
func GetTestFilePath() string {
	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	// return os.Getenv("COMMON_PATH") + "golang-studies/06-parquet-converter/internal/pkg/data/test/"
	return "/home/pedro/dev/golang-studies/06-parquet-converter/internal/pkg/data/test/"
}
