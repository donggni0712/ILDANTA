# ILDANTA

<p align="center">
	<a href="https://https://github.com/donggni0712/ILDANTA/search?l=JavaScript&type=code"><img alt="GitHub language count" src="https://img.shields.io/github/languages/count/donggni0712/ILDANTA"></a>
	<a href="https://github.com/donggni0712/ILDANTA/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/donggni0712/ILDANTA?color=success"></a>
	<a href="https://github.com/donggni0712/ILDANTA/stargazers"><img alt="GitHub stars" src="https://img.shields.io/github/stars/donggni0712/ILDANTA"></a>
	<a href="https://github.com/donggni0712/ILDANTA/network/members"><img alt="GitHub forks" src="https://img.shields.io/github/forks/donggni0712/ILDANTA"></a>
	<a href="https://github.com/donggni0712/ILDANTA/blob/master/LICENSE"><img alt="GitHub license" src="https://img.shields.io/github/license/donggni0712/ILDANTA"></a><br>
    <a href="https://www.youtube.com/watch?v=qoCpZwnwJdk"><img src="https://img.shields.io/youtube/likes/qoCpZwnwJdk?style=social"></a>
  </p>

![LOGO](./DOC/img/LOGO.png)
급하게 대중교통을 탈 땐, 일단 타!

## Mock-up

![MockUp](./DOC/img/ILDANTA_UI.gif)

## Flow Chart

![FlowChart](./DOC/img/ILDANTA_FlowChart.png)

## Service

**대중교통 길찾기 기능을 이용할 수 있고, 어떤 버스를 탄 이후의 환승정보를 보기 간편하게 출력해준다.**

예를 들어, 7000번을 이미 탔다면 7000번을 탄 후 어디서 내려서 어떤 버스로 환승하는 지를 보여준다.
| 메인화면 | 길찾기 | 경로 출력1 | 경로 출력2 |
| :--------------------------: | :-------------------------------: | :-------------------------------: | :-------------------------------: |
| ![main](./DOC/img/Main1.png) | ![Search1](./DOC/img/Search1.png) | ![Search2](./DOC/img/Search2.png) | ![Search3](./DOC/img/Search3.png) |

## RestAPI

##### [API REFERENCE 참조하기](./BE/README.md#api-reference)

- **Search** : [POST] /Search
- **Choose** : [POST] /Search/Choose

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

- [ ] 현재 데이터가 [ODsay => ResultDomain => restDomain] 순으로 정제됨. 이를 단순화? 시키면 좋을지
- [ ] DB연결
- [ ] FirstRoute와 SubRoute의 rest API response 값이 다른데 이를 통일 시킬 필요가 있을까?
