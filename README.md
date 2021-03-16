# Go-Bit! Sample Code
* [shieldnet/gobit](https://github.com/shieldnet/gobit)를 이용한 Upbit 코인 자동 매매 프로그램 샘플 코드입니다.

## How to use?
* `go`를 설치합니다. [Download and install go](https://golang.org/doc/install)
* `${GOPATH}/src/github.com/shieldnet` 디렉토리를 만들어줍니다.
  * fork나 clone해서 사용하시는 경우, `${GOPATH}/src/github.com/<your_name>` 디렉토리를 만들어줍니다.
* `cd gobit-sample`을 통해 gobit 폴더로 이동해줍니다.
* `go get` 명령어를 실행해 `go.mod`의 패키지를 다운로드해줍니다.
  * go mod require의 `github.com/shieldnet/gobit`의 버전이 다를 수 있습니다. 참고하시기 바랍니다.
* [Upbit OpenAPI 관리](https://upbit.com/mypage/open_api_management) 페이지에서 Open API Key를 발급받습니다.
  * SecretKey는 발급받은 직후가 아니면 다시 볼 수 없습니다. 주의하시기 바랍니다.
* `main.go`폴더의 keys에 발급받은 SecretKey와 AccessKey를 입력합니다.
  * **주의 : 발급받은 key를 github이나 다른 웹 상의 다른 공개된 장소에 업로드 하지 않도록 주의합니다.**
* 전략을 본인의 취향에 맞게 수정합니다.
  * 참고) 현재 코드에 작성된 전략은 다음과 같습니다.
    * 2% 손해 보면 손절
    * 최근으로부터 5개 봉을 보고, 그 중 현재의 봉이 최저가이면 매수.
    * 매수를 했을 때, 최근으로부터 5개 봉을 보고 현재의 봉이 최고가이면 매도.
    * 반복
* `go run main.go`로 전략을 실행합니다.

## 사용 결과 화면 (예시)
![image](https://user-images.githubusercontent.com/9548599/111020890-4e1e1380-840c-11eb-8c59-69141c5f7c9b.png)

## Caution(주의사항)
* 본 자동 매매 프로그램을 이용해서 손해가 발생하더라도, 저는 책임지지 않습니다.
* 본 자동 매매 프로그램을 이용하는 순간 `주의사항`에 동의한 것으로 간주합니다.
* `gobit` API를 제외한 샘플 코드는 자유롭게 변형해 사용하셔도 좋습니다.
* 본 소프트웨어의 상업적, 다수를 대상으로 한 이윤 창출 목적의 이용을 금지합니다. 단, 개인 사용자의 개인 투자 목적으로는 사용 가능합니다. 

## Contribution
* 기본적으로, 개인 프로젝트이기 때문에 Contribution은 받지 않습니다만, 버그나 기능 추가에 관한 Issue는 환영하고 있습니다.
* 질문이나 버그는 이슈로 달아주시거나, 개발자 이메일 (atez.dev@gmail.com)으로 보내주시기 바랍니다.

## FAQ
### 거래가 너무 느려요!!
* 기본적으로 upbit에는 API call 횟수 제한이 있기 때문에 각 단계 별 요청을 보내는데 시간이 조금 걸리게 해두었습니다. `strategy/condition.go` 에서 ` IntervalTime`을 바꿔주시면 속도를 바꾸실 수 있습니다.

