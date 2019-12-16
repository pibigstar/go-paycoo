package paycoo

import (
	"testing"
)

var (
	sign       = "cRieysfAY04h+/3ZWgU6XTIgxGGM8/UCs/i6kFpyqj1T+bhx5WmvqDrpi5fGdVeopDPOOwBLNHs2j8UlukttYMDsudlMin4Zap8QCuXNxRg1/CtbzOyJnWOYveZf9+w5Y160gLzRO4PgYVOuf+j6mKCw8HDPjf++nCwrc6PkalCqhazshM+85GTYcAgDvU4k5nRKWWH6dPP5MYuTkMNDrhTlb5vcYkjsvLSO+YJ9RCoqBG5ubWx9aExiKI3e7Cx3nvemFK5fxMglZ6hNVk9OJcKJVp7v1ryj3vOuqXCZfMXCLDE+1LnevaMBPrtg9wOehC8HEU1V/XcEND3OAQG0oA=="
	src        = `code=0&data=[{"bill_file_url":"https://jf-ibp.oss-cn-beijing.aliyuncs.com/accountingProd/bill/d040aeaf23a44dcd/3703BC3BCBF10F7DFAC0284B33FB139A.txt"}]&msg=success&psn=11191409431255382302&total=1`
	returnSign = "fsYx3zMQwSViD8QYxTAZhayugFDYSF3MnVwoxXod6jEcPF3QMwJpQr6beJqdeTT62UpbXncyw9LygI8X592nR+KCYP9OA1Ouay/3/8tVYZlOnZAmmStY6ybZ55iX5fIm71ab9YXCvZQHXoftPKclack1oKKS8Gep5FGV+JbnUuHYTkd2FRUWWgXUUxM4pRU153P058eFjm59hxVD9XnC5457RF+DZupDgzAI9AwuvJNTu31CemsH3qlSxv9Yw40a0HtJ/Sai4RnYQBbi+IJKsR4TrQTHertBU+BdkYxW5Di2790UJYnfKm2Jd7jNCnEoKb3OmQQhrnsKxaddAFsBpA=="
)

func TestRsaSign(t *testing.T) {
	key, err := ParsePrivateKey(privateKey)
	if err != nil {
		t.Error(err)
	}
	sign, err = sha256WithRsaWithBase64([]byte(src), key)
	if err != nil {
		t.Error(err)
	}
	t.Log(sign)

	err = VerifySign([]byte(src), sign, appPublicKey)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("verify success!")
	}

}

func TestVerifySign(t *testing.T) {
	err := VerifySign([]byte(src), returnSign, publicKey)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("verify success!")
	}
}
