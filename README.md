# img2svg
이미지 파일을 svg 파일로 변경하는 코드입니다.
macOS에서 제가 사용하기 위해서 만든 코드입니다.
imagemagick/convert, potrace 명령어가 필요합니다.
potrace는 bmp만 잘 인식하여 매번 bmp로 컨버팅후 potrace 명령을 사용하는것이 불편하여 제작하게 되었습니다.

#### Pre Install
```
brew install imagemagick
brew install potrace
```

#### 사용법
- ~/test.png 파일을 이용해서 ~/test.svg 파일을 생성
```
$ img2svg -src ~/test.png
```

#### Install from golang
```
$ go get -u github.com/lazypic/img2svg
```
