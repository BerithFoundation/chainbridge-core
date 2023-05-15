// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package keystore

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ChainSafe/chainbridge-core/crypto/secp256k1"
	"github.com/ChainSafe/chainbridge-core/crypto/sr25519"
)

func TestEncryptAndDecrypt(t *testing.T) {
	password := []byte("noot")
	msg := []byte("helloworld")

	ciphertext, err := Encrypt(msg, password)
	if err != nil {
		t.Fatal(err)
	}

	res, err := Decrypt(ciphertext, password)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(msg, res) {
		t.Fatalf("Fail to decrypt: got %x expected %x", res, msg)
	}
}

func TestEncryptAndDecryptKeypair(t *testing.T) {
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	kp, err := secp256k1.NewKeypairFromPrivateKey(buf)
	if err != nil {
		t.Fatal(err)
	}

	password := []byte("noot")

	data, err := EncryptKeypair(kp, password)
	if err != nil {
		t.Fatal(err)
	}

	res, err := DecryptKeypair(kp.PublicKey(), data, password, "secp256k1")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(kp, res) {
		t.Fatalf("Fail: got %#v expected %#v", res, kp)
	}
}

func createTestFile(t *testing.T) (*os.File, string) {
	filename := "./test_key"

	fp, err := filepath.Abs(filename)
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Create(fp)
	if err != nil {
		t.Fatal(err)
	}

	return file, fp
}

func TestEncryptAndDecryptFromFile_Secp256k1(t *testing.T) {
	password := []byte("noot")
	file, fp := createTestFile(t)
	defer os.Remove(fp)

	kp, err := secp256k1.GenerateKeypair()

	if err != nil {
		t.Fatal(err)
	}

	err = EncryptAndWriteToFile(file, kp, password)
	if err != nil {
		t.Fatal(err)
	}

	res, err := ReadFromFileAndDecrypt(fp, password, "secp256k1")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(kp.Encode(), res.Encode()) {
		t.Fatalf("Fail: got %#v expected %#v", res, kp)
	}
}

func TestEncryptAndDecryptFromFile_Sr25519(t *testing.T) {
	password := []byte("ansermino")
	file, fp := createTestFile(t)
	defer os.Remove(fp)

	kp, err := sr25519.NewKeypairFromSeed("//seed", "substrate")
	if err != nil {
		t.Fatal(err)
	}

	err = EncryptAndWriteToFile(file, kp, password)
	if err != nil {
		t.Fatal(err)
	}

	res, err := ReadFromFileAndDecrypt(fp, password, "sr25519")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(kp.Encode(), res.Encode()) {
		t.Fatalf("Fail: got %#v expected %#v", res, kp)
	}
}

func TestDecryptIncorrectType(t *testing.T) {
	password := []byte("123")

	kp, err := ReadFromFileAndDecrypt("./test_key", password, "secp256k1")
	if err != nil {
		t.Fatal("mismatch error")
	}
	t.Log(kp.Encode())
	t.Log(hex.EncodeToString([]byte{109, 246, 186, 111, 86, 245, 123, 175, 124, 107, 221, 223, 225, 247, 123, 243, 157, 28, 107, 167, 90, 119, 198, 157, 119, 222, 26, 219, 126, 53, 121, 182, 221, 119, 166, 158, 111, 199, 189, 119, 190, 121, 223, 110, 157, 121, 167, 157}))

}

func TestDecrypt(t *testing.T) {
	password := []byte("123")
	cipher := []byte{109, 246, 186, 111, 86, 245, 123, 175, 124, 107, 221, 223, 225, 247, 123, 243, 157, 28, 107, 167, 90, 119, 198, 157, 119, 222, 26, 219, 126, 53, 121, 182, 221, 119, 166, 158, 111, 199, 189, 119, 190, 121, 223, 110, 157, 121, 167, 157}
	b, err := Decrypt(cipher, password)
	if err != nil {
		t.Error(err, "failed")
	}
	t.Log(string(b))

	//TODO : Load Privatekey
}
