package hashid

import (
	"fmt"
	"testing"

	"github.com/speps/go-hashids"
	"github.com/stretchr/testify/assert"
)

func Test_InitIDHasher(t *testing.T) {
	t.Parallel()

	t.Run("singleton instance exists", func(t *testing.T) {
		backup := singleton
		defer func() {
			singleton = backup
		}()

		singleton = &idHasher{}
		minLength := 16
		salt := "secret_salt"

		gotErr := InitIDHasher(minLength, salt)
		assert.Nil(t, gotErr)
	})

	t.Run("success", func(t *testing.T) {
		backup := singleton
		defer func() {
			singleton = backup
		}()

		singleton = nil
		minLength := 16
		salt := "secret_salt"

		gotErr := InitIDHasher(minLength, salt)

		assert.Nil(t, gotErr)
	})
}

func Test_GetIDHasher(t *testing.T) {
	t.Parallel()

	t.Run("get nil", func(t *testing.T) {
		backup := singleton
		defer func() {
			singleton = backup
		}()

		singleton = nil

		got := GetIDHasher()

		assert.Nil(t, got)
	})

	t.Run("get singleton", func(t *testing.T) {
		backup := singleton
		defer func() {
			singleton = backup
		}()

		singleton = &idHasher{}

		got := GetIDHasher()

		assert.NotNil(t, got)
	})
}

func Test_idHasher_Encode(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		id := 0
		hashID, _ := hashids.NewWithData(&hashids.HashIDData{
			Alphabet:  hashids.DefaultAlphabet,
			MinLength: 16,
			Salt:      "salt",
		})
		hashed, _ := hashID.Encode([]int{id})
		idHasher := &idHasher{
			hashID: hashID,
		}

		got := idHasher.Encode(id)

		assert.Equal(t, hashed, got)
	})
}

func Test_idHasher_Decode(t *testing.T) {
	t.Parallel()

	t.Run("incorrect hash id format", func(t *testing.T) {
		hashID, _ := hashids.NewWithData(&hashids.HashIDData{
			Alphabet:  hashids.DefaultAlphabet,
			MinLength: 16,
			Salt:      "salt",
		})
		hashed, _ := hashID.Encode([]int{1, 2, 3})
		idHasher := &idHasher{
			hashID: hashID,
		}

		wantInteger := 0
		wantErr := fmt.Errorf("incorrect hash id format")
		gotInteger, gotErr := idHasher.Decode(hashed)

		assert.Equal(t, wantErr, gotErr)
		assert.Equal(t, wantInteger, gotInteger)
	})

	t.Run("success", func(t *testing.T) {
		hashID, _ := hashids.NewWithData(&hashids.HashIDData{
			Alphabet:  hashids.DefaultAlphabet,
			MinLength: 16,
			Salt:      "salt",
		})
		hashed, _ := hashID.Encode([]int{1})
		idHasher := &idHasher{
			hashID: hashID,
		}

		wantInteger := 1
		gotInteger, gotErr := idHasher.Decode(hashed)

		assert.Nil(t, gotErr)
		assert.Equal(t, wantInteger, gotInteger)
	})
}
