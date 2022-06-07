# API REFERENCE

백엔드에서 프론트엔드에게 제공하는 API는 다음과 같습니다.<br/>
백엔드는 port 3001에서 실행됩니다.<br/>
go설치 후 `go run main.go`로 서버를 실행시킬 수 있습니다.

# Search

출발지로부터 목적지까지의 경로를 찾는다.
| method | request URL | format |
| :----: | :--------------------------: | :----: |
| POST | http://localhost:3001/Search | json |

### request body

| key | valueType | requirement |  describtion   |
| :-: | :-------: | :---------: | :------------: |
| sx  |  string   |      Y      | 출발지의 x좌표 |
| sy  |  string   |      Y      | 출발지의 y좌표 |
| ex  |  string   |      Y      | 목적지의 x좌표 |
| ey  |  string   |      Y      | 목적지의 y좌표 |

## response body

|     key     | value  |            describtion             |
| :---------: | :----: | :--------------------------------: |
|  whereOns   | Array  |          교통 정보 리스트          |
|   whereOn   | string |            정류장 이름             |
|   whatOns   | Array  | 해당 정류장에서의 교통 정보 리스트 |
|   whatOn    | string |          버스/지하철 이름          |
| transferNum | string |              환승 수               |
|  totalTime  | string |             소요 시간              |

## example

- request body

```json
{
  "sx": "127.08186574229312",
  "sy": "37.23993898645113",
  "ex": "127.05981200975921",
  "ey": "37.28556112210226"
}
```

- response body

```json
{
  "whereOns": [
    {
      "whereOn": "서천2차아이파크",
      "whatOns": [
        {
          "whatOn": "8",
          "transferNum": "0번~3번",
          "totalTime": "63분"
        }
      ]
    },
    {
      "whereOn": "사색의광장",
      "whatOns": [
        {
          "whatOn": "1112",
          "transferNum": "0번~3번",
          "totalTime": "60분"
        },
        {
          "whatOn": "5100",
          "transferNum": "3번",
          "totalTime": "60분"
        },
        {
          "whatOn": "7000",
          "transferNum": "3번",
          "totalTime": "60분"
        },
        {
          "whatOn": "9",
          "transferNum": "0번~3번",
          "totalTime": "50분"
        }
      ]
    },
    {
      "whereOn": "경희대학교",
      "whatOns": [
        {
          "whatOn": "7-2",
          "transferNum": "0번~1번",
          "totalTime": "88분"
        }
      ]
    }
  ]
}
```

# Choose

반환된 경로에서 이용할 교통수단을 고른 후의 교통수단에 대해 경로를 찾는다.
| method | request URL | format |
| :----: | :---------------------------------: | :----: |
| POST | http://localhost:3001/Search/Choose | json |

### request body

|     key     | valueType | requirement |          describtion           |
| :---------: | :-------: | :---------: | :----------------------------: |
| coordinate  | 확장 노드 |      Y      | 좌표값을 가지고 있는 확장 노드 |
|     sx      |  string   |      Y      |         출발지의 x좌표         |
|     sy      |  string   |      Y      |         출발지의 y좌표         |
|     ex      |  string   |      Y      |         목적지의 x좌표         |
|     ey      |  string   |      Y      |         목적지의 y좌표         |
| firstChoice | 확장 노드 |      Y      |       교통수단 확장노드        |
|   whereOn   |  string   |      Y      |       탑승한 정류장 이름       |
|   whatOn    |  string   |      Y      |      탑승한 교통수단 이름      |
|   choices   |   Array   |      N      |  이용한 교통수단 정보 리스트   |
|  whereOff   |  string   |      N      |       하차한 정류장 이름       |
|   whereOn   |  string   |      N      |       탑승한 정류장 이름       |
|   whatOn    |  string   |      N      |      탑승한 교통수단 이름      |

## response body

|     key     | value  |            describtion             |
| :---------: | :----: | :--------------------------------: |
| whatTookOn  | string |          버스/지하철 이름          |
| whereTookOn | string |            정류장 이름             |
|  whereOffs  | Array  |        하차할 정류장 리스트        |
|  whereOff   | string |            정류장 이름             |
|  whereOns   | Array  |          교통 정보 리스트          |
|   whereOn   | string |            정류장 이름             |
|   whatOns   | Array  | 해당 정류장에서의 교통 정보 리스트 |
|   whatOn    | string |          버스/지하철 이름          |
| transferNum | string |              환승 수               |
|  totalTime  | string |             소요 시간              |

## example

- request body

```json
{
  "coordinate": {
    "sx": "127.08186574229312",
    "sy": "37.23993898645113",
    "ex": "127.05981200975921",
    "ey": "37.28556112210226"
  },
  "firstChoice": {
    "whereOn": "사색의광장",
    "whatOn": "7000"
  }
}
```

- response body

```json
{
  "whatTookOn": "7000",
  "whereTookOn": "사색의광장",
  "whereOffs": [
    {
      "whereOff": "살구골현대아파트.영통역4번출구",
      "whereOns": [
        {
          "whereOn": "영통",
          "whatOns": [
            {
              "whatOn": "수도권 수인.분당선",
              "transferNum": "",
              "totalTime": ""
            }
          ]
        }
      ]
    },
    {
      "whereOff": "아주대.아주대학교병원",
      "whereOns": [
        {
          "whereOn": "아주대.아주대학교병원",
          "whatOns": [
            {
              "whatOn": "81",
              "transferNum": "",
              "totalTime": ""
            },
            {
              "whatOn": "13-4",
              "transferNum": "",
              "totalTime": ""
            },
            {
              "whatOn": "20",
              "transferNum": "",
              "totalTime": ""
            },
            {
              "whatOn": "7",
              "transferNum": "",
              "totalTime": ""
            }
          ]
        }
      ]
    },
    {
      "whereOff": "KT동수원지사",
      "whereOns": [
        {
          "whereOn": "KT동수원지사",
          "whatOns": [
            {
              "whatOn": "720-3",
              "transferNum": "",
              "totalTime": ""
            }
          ]
        }
      ]
    },
    {
      "whereOff": "영통역",
      "whereOns": [
        {
          "whereOn": "영통역",
          "whatOns": [
            {
              "whatOn": "2-1",
              "transferNum": "",
              "totalTime": ""
            }
          ]
        }
      ]
    }
  ]
}
```
