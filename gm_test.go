package mogrify

import (
  "io/ioutil"
  "log"
  "os"
  "testing"
)

func TestOpenExisting(t *testing.T) {
  img := Open("./assets/image.jpg")
  if img == nil {
    t.Fail()
  }
  img.Destroy()
}

func TestOpenNonExisting(t *testing.T) {
  if Open("./assets/image_does_not_exist.jpg") != nil {
    t.Fail()
  }
}

func TestResizeSuccess(t *testing.T) {
  img := Open("./assets/image.jpg")

  if img == nil {
    t.Fail()
    return
  }

  defer img.Destroy()

  status := img.Resize(50, 50)
  if status != nil {
    log.Printf("resize failed %s", status)
    t.Fail()
  }
}

func TestResizeFailure(t *testing.T) {
  img := Open("./assets/image.jpg")

  if img == nil {
    t.Fail()
    return
  }

  defer img.Destroy()

  status := img.Resize(0, 50)
  if status == nil {
    t.Fail()
  }
}

func TestSaveToSuccess(t *testing.T) {
  img := Open("./assets/image.jpg")

  if img == nil {
    t.Fail()
    return
  }

  defer img.Destroy()

  res := img.SaveFile("/tmp/img.jpg")
  if res != nil {
    t.Fail()
  }
}

func TestSaveToFailure(t *testing.T) {
  img := Open("./assets/image.jpg")

  if img == nil {
    t.Fail()
    return
  }

  defer img.Destroy()

  res := img.SaveFile("/dgksjogdsksdgsdkgsd;lfsd-does-not-exist/img.jpg")
  if res == nil {
    t.Fail()
  }
}

func TestOpenBlopSuccess(t *testing.T) {
  bytes, _ := ioutil.ReadFile("./assets/image.jpg")

  img := NewImage()
  res := img.OpenBlob(bytes)

  if res != nil {
    t.Fail()
  }

  img.Destroy()
}

func TestOpenBlopFailure(t *testing.T) {

  img := NewImage()
  res := img.OpenBlob([]byte{'a'})

  if res == nil {
    t.Fail()
  }

  res = img.OpenBlob([]byte{})

  if res == nil {
    t.Fail()
  }
}

func TestSaveToBlob(t *testing.T) {
  img := Open("./assets/image.jpg")

  fp, err := os.Create("/tmp/img3.jpg")
  if err != nil {
    t.Fail()
  }

  defer fp.Close()

  n, err := img.Write(fp)

  if err != nil {
    t.Fail()
  }

  log.Printf("%d", n)

}
