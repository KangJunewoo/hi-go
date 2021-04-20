# hi-go
노마드코더와 함께하는 Go 입문

## 기본문법
코드 보면 바로 알 수 있을 것이다.
channels랑 go routines는 추후에 다룰 예정.
하이레벨 + 로우레벨을 다 아우르는 언어같다.


## working directory is not part of a module 에러뜰때
```bash
$ go mod init
```
해주면 됨.

## 고루틴
병렬처리를 위해 사용됨.

## channel
고루틴 - 메인함수가 정보를 주고받기 위해 쓰는 통로.

와.. 파이썬이랑 비교가 안되게 빠르다.

## goquery
고의 뷰티풀숩 같은거.  
```
$ go get github.com/PuerkitoBio/goquery
```
로 다운로드 가능.