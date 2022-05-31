import React, { useState } from "react";
import Map from "./map.js";

const {kakao} = window;

const SearchPlace = () => {
  const [inputStart, setInputStart] = useState("");
  const [inputEnd, setInputEnd] = useState("");
  const [place, setPlace] = useState("");
  const [isStartSubmit, setIsStartSubmit] = useState(true);
  const [start,setStart] = useState({name: "", x:0,y:0});
  const [end,setEnd] = useState({name: "", x:0,y:0});
  const [isSearched,setIsSearched] = useState(false);
  const [result, setResult] = useState({});

  const onChangeStart = (e) => {
    setInputStart(e.target.value);
  };
  const onChangeEnd = (e) => {
    setInputEnd(e.target.value);
  };
  const handleStartSubmit = (e) => {
    setIsStartSubmit(true)
    e.preventDefault();
    setPlace(inputStart);
    setInputStart("");
  };

   const handleEndSubmit = (e) => {
    setIsStartSubmit(false)
    e.preventDefault();
    setPlace(inputEnd);
    setInputEnd("");
  };

  const handleSearch = () =>{
    console.log('loading...')
    setIsSearched(true)
    fetch(`http://localhost:3001/Search/${start.x}&${start.y}&${end.x}&${end.y}`)
    .then((respons)=>respons.json())
    .then((res)=>{
      console.log(res);
      setResult(res);
    })
  }

  function ClickList(item){
    const input = {
      name : item.place_name,
      x : item.x,
      y : item.y
    }
    if(isStartSubmit){
      setStart(input)
    }
    if(!isStartSubmit){
      setEnd(input)
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

      <Map searchPlace={place} ClickList={ClickList} isSearched={isSearched}/>

      <div className='startPlace'>출발지 : {start.name}</div>
      <div className='startPlace'>도착지 : {end.name}</div>
    </>
  );
};

export default SearchPlace;