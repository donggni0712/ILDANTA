# ILDANTA

![LOGO](./DOC/img/LOGO.png)

## Mock-up

![MockUp](./DOC/img/ILDANTA_UI.gif)

## RestAPI

- [GET] "/development/RawData/{sx}&{sy}&{ex}&{ey}"
  - ODsayAPI를 정제하기만 한 날 것의 데이터 보기
- [GET] "/Search/{sx}&{sy}&{ex}&{ey}"
  - 출발지에서 도착지까지 가는 경로 출력
- [GET] "/Search/ChooseTakeOn/{whereOn}&{whatOn}" body : {json}
  - Search에서 어디서 무엇을 탈 지 선택하면 해당 경로의 하위 경로 출력
  - request body에는 /Search의 response body 입력
- [GET] "Search/ChooseTakeOffOn/{whereOff}&{whereOn2}&{whatOn2}" body : {json}
  - Search/ChooseTakeOn에서 선택한 경로의 하위 경로 출력
  - request body에는 /Search/ChooseTakeOn의 response body 입력

## ToDo!!

- [x] MaxTransferNum, MinTransferNum 반영(21.11.30)
- [x] MaxTotalTime, MinTotalTime 반영(21.11.30)
- [x] Figma 참고해서 출력함수 생성(21.11.30)
- [x] 출력에서 '여기서' 통합, '이거타면' 통합.(21.12.01)
- [x] Code Refactor(22.05.30 - 아직 부족)
- [ ] Code Refactor More
- [x] 현재 재귀적인 struct가 구현되지 않음. 재귀적으로 구현 (21.12.03)
- [x] RestAPI 구조 개선 필요(body값을 넣고 읽어오는 것 구현하기)(22.05.28)
- [x] RestAPI의 재귀적 호출 구현(22.05.30 - 굳이 구현할 필요 없을 것으로 예상)
