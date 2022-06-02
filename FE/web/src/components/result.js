import React from "react";
import './result.css'
import { useState } from "react";

const Result = ({response,subPage,ClickFirstPath,isFirst}) =>{

    if(isFirst==true){
    return <div className="result">
        <a>여기서</a> <a className="whereOn">이거타</a>
        {response.whereOns.map((elWhere)=>{
        return <div>
            {elWhere.whereOn}

            {elWhere.whatOns.map((el)=>{
                return <div className="whereOn" onClick={()=>ClickFirstPath(elWhere.whereOn,el.whatOn)}>
                    {el.whatOn}
                    </div>
            })}

        </div>
    })}
    </div>
    }
    return<div className="result">
        <a>여기서 내려서</a>  <a className="whereOn">여기서</a><a className="whereOn">이거타</a>
        {subPage.whereOffs.map((elWhereoff)=>{
        return <div>
            {elWhereoff.whereOff}
            {elWhereoff.whereOns.map((elWhereOn)=>{
                return <div className="whereOn">
                    <a>{elWhereOn.whereOn}</a>
                    {elWhereOn.whatOns.map((el)=>{
                        return<div className="whereOn">{el.whatOn}</div>
                    })}
                    </div>
            })}

        </div>
    })}
    </div>
}

export default Result;