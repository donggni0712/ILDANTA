import React, { useState } from "react";
import Map from "./map.js";
import Result from "./result.js";

const { kakao } = window;

const SearchPlace = () => {
  const [choices, setChoices] = useState([]);
  const [firstChoice, setFirstChoice] = useState({
    whereOn: "",
    whatOn: "",
  });
  const [inputStart, setInputStart] = useState("");
  const [inputEnd, setInputEnd] = useState("");
  const [place, setPlace] = useState("");
  const [isStartSubmit, setIsStartSubmit] = useState(true);
  const [start, setStart] = useState({ name: "", x: 0, y: 0 });
  const [end, setEnd] = useState({ name: "", x: 0, y: 0 });
  const [isSearched, setIsSearched] = useState(false);
  const [isFirst, setIsFirst] = useState(true);
  const [response, setResponse] = useState({
    whereOns: [
      {
        whereOn: null,
        whatOns: [
          {
            whatOn: null,
            transferNum: null,
            totalTime: null,
          },
        ],
      },
    ],
  });

  const [subPage, setSubPage] = useState({
    whatTookOn: null,
    whereTookOn: null,
    whereOffs: [
      {
        whereOff: null,
        isFinal: 0,
        whereOns: [
          {
            whereOn: null,
            whatOns: [
              {
                whatOn: null,
                transferNum: null,
                totalTime: null,
              },
            ],
          },
        ],
      },
    ],
  });
  const onChangeStart = (e) => {
    setInputStart(e.target.value);
  };
  const onChangeEnd = (e) => {
    setInputEnd(e.target.value);
  };
  const handleStartSubmit = (e) => {
    setIsStartSubmit(true);
    e.preventDefault();
    setPlace(inputStart);
    setInputStart("");
  };

  const handleEndSubmit = (e) => {
    setIsStartSubmit(false);
    e.preventDefault();
    setPlace(inputEnd);
    setInputEnd("");
  };

  const handleSearch = () => {
    // console.log('loading...')
    setIsSearched(true);
    const coordinate = JSON.stringify({
      sx: `${start.x}`,
      sy: `${start.y}`,
      ex: `${end.x}`,
      ey: `${end.y}`,
    });
    fetch(`http://localhost:3001/Search`, {
      method: "POST",
      body: coordinate,
    })
      .then((respons) => respons.json())
      .then((res) => {
        setIsFirst(true);
        setResponse(res);
      });
  };

  const ClickFirstPath = (_whereOn, _whatOn) => {
    console.log("loading...");
    setIsSearched(true);
    const coordinate = {
      sx: `${start.x}`,
      sy: `${start.y}`,
      ex: `${end.x}`,
      ey: `${end.y}`,
    };

    setFirstChoice({
      whereOn: _whereOn,
      whatOn: _whatOn,
    });

    const requestBody = JSON.stringify({
      coordinate: coordinate,
      firstChoice: {
        whereOn: _whereOn,
        whatOn: _whatOn,
      },
    });
    console.log(_whereOn, _whatOn);
    console.log(requestBody);
    fetch(`http://localhost:3001/Search/Choose`, {
      method: "POST",
      body: requestBody,
    })
      .then((respons) => respons.json())
      .then((res) => {
        console.log("here");
        console.log(res);
        setIsFirst(false);
        setSubPage(res);
      });
  };

  const ClickPath = (_whereOff, _whereOn, _whatOn) => {
    console.log("loading...");
    setIsSearched(true);
    const coordinate = {
      sx: `${start.x}`,
      sy: `${start.y}`,
      ex: `${end.x}`,
      ey: `${end.y}`,
    };
    const requestChoices = {
      whereOff: _whereOff,
      whereOn: _whereOn,
      whatOn: _whatOn,
    };
    const requestBody = JSON.stringify({
      coordinate: coordinate,
      firstChoice: firstChoice,
      choices: [...choices, requestChoices],
    });

    setChoices([...choices, requestChoices]);
    console.log(_whereOn, _whatOn);
    console.log(requestBody);
    fetch(`http://localhost:3001/Search/Choose`, {
      method: "POST",
      body: requestBody,
    })
      .then((respons) => respons.json())
      .then((res) => {
        console.log("here");
        console.log(res);
        setIsFirst(false);
        setSubPage(res);
      });
  };

  function ClickList(item) {
    const input = {
      name: item.place_name,
      x: item.x,
      y: item.y,
    };
    if (isStartSubmit) {
      setStart(input);
    }
    if (!isStartSubmit) {
      setEnd(input);
    }
  }

  return (
    <>
      <form className="inputForm" onSubmit={handleStartSubmit}>
        <label>출발지 : </label>
        <input
          placeholder="Search Place..."
          onChange={onChangeStart}
          value={inputStart}
        />
        <button type="submit">검색</button>
      </form>
      <form className="inputForm" onSubmit={handleEndSubmit}>
        <label>도착지 : </label>
        <input
          placeholder="Search Place..."
          onChange={onChangeEnd}
          value={inputEnd}
        />
        <button type="submit">검색</button>
      </form>
      <button onClick={handleSearch}>경로 검색</button>

      <Map searchPlace={place} ClickList={ClickList} isSearched={isSearched} />

      <div className="startPlace">출발지 : {start.name}</div>
      <div className="startPlace">도착지 : {end.name}</div>

      <Result
        response={response}
        subPage={subPage}
        ClickFirstPath={ClickFirstPath}
        ClickPath={ClickPath}
        isFirst={isFirst}
      />
    </>
  );
};

export default SearchPlace;
