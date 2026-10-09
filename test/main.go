package main

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

// **Описание**: Реализуйте функцию для расшифровки зашифрованных данных с использованием AES-GCM
//
// **Входные данные**: Зашифрованные данные в виде []byte (встроенные в код как переменная)
//
// **Выходные данные**: Расшифрованные данные в виде []byte или ошибка при неудачной расшифровке
//
// **Ограничения**: Используйте только стандартные пакеты Go (crypto/aes, crypto/cipher, errors)
//
// **Примеры**:
// Input: []byte{...зашифрованные данные с nonce в начале...}
// Output: []byte("Hello, World!")
//
// Input: []byte{...другие зашифрованные данные...}
// Output: []byte("Secret message")

type Encrypter struct {
	key []byte
}

func (e *Encrypter) Decrypt(data []byte) ([]byte, error) {
	// Ваш код здесь
	if len(data) == 0 {
		return nil, errors.New("Should be encrypted data")
	}
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

func main() {
	// Пример зашифрованных данных для тестирования
	encryptedData := []byte{ /* здесь будут зашифрованные данные */ }

	encrypter := &Encrypter{key: []byte("your-32-byte-key-here-for-aes256")}
	// Ваш код здесь
	decrypted, err := encrypter.Decrypt(encryptedData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Расшифрованные данные:", string(decrypted))
}
