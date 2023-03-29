'use client';
import './user.css'
import Image from 'next/image'
import placeholder from '../../public/placeholder.jpg'
import { useSearchParams } from 'next/navigation';
import {useState} from 'react';
var XMLHttpRequest = require('xhr2');
var response2;
var render;

function getData(param) {
  let url = "http://127.0.0.1:3000/api/userinfo/" + param;
  let request = new XMLHttpRequest()
  request.open("GET",url);
  request.send();
  request.onload = () => {
  if(request.status === 200){
  response2 = (JSON.parse(request.response))
  } else {
  console.log("Page not found")// if link is broken, output will be page not found
  }
  }
  }

  async function getData2(param) {
    let url = "http://127.0.0.1:3000/api/userinfo/" + param;;
    let response = await
    fetch(url)
    let data = await response.json()
    return data;
    }

function displayData(data){
  return (
    <div id="user_info_container">
      <p className="info_tittle">Puntos adquiridos</p>
      <div className="user_info">
          <i class="fa-sharp fa-solid fa-coins price_icon_3"></i>
          <p id ="coins_user">{data.coins}</p>
      </div>
      <p className="info_tittle">Fila Actual</p>
      <div id="image_container"> 
      <Image className="atraccion_image"src={placeholder}/>
      </div>
      <p id="user_turn_info">Atraccion: Turno {data.turn}</p>
    </div>
  )
}


export default function Usuarios(props) {
  const[data,setData] = useState("")
  const handle = async (event) => {
    const id = (event.target.previousElementSibling.value);
    const data = await (getData2(id))
    console.log(data);
    if(data.code != 404){
    setData(displayData(data))
    } 
  };
    return (
      <main>
        <header>
          <h1 id="header-text"> Usuarios </h1>
        </header>
        <div id="user-form">
          <label htmlFor="user-input"> <p id="input-label">Ingrese su ID</p></label>
          <input id="user-input"type="number" name="id" required/>
          <button onClick={handle} value="holaa" id="user-submit"> Consultar</button>
        </div>
        <p>{data}</p>
      </main>
    )
  }
  