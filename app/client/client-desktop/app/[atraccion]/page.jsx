import Image from 'next/image'
import Link from 'next/link'
import placeholder from '../../public/placeholder.jpg'
import './atraccion.css'
export const metadata = {
    title: 'Usuarios',
  }

const atracciones = [
  {
    "capacity": "50",
    "description": "Fire",
    "duration": "60",
    "name": "atraccion",
    "nextTurn": 50,
  },
  {
    "capacity": "20",
    "description": "water",
    "duration": "6",
    "name": "atraccion extrema",
    "nextTurn": 501,

  },
  {
    "capacity": "1",
    "description": "Earth",
    "duration": "1",
    "name": "atraccion zzzzz",
    "nextTurn": 10,
  }
]


async function getData() {
  let response = await
  fetch("http://127.0.0.1:3000/api/attractioninfo/Divertido")
  let data = await response.json()
  return data;
  }


export default async function Atraccion(param) {
    const atraccion = await getData();
    //const atraccion = atracciones[(param.params.atraccion)] 
    const elements = []
    for (let i = atraccion.nextTurn; i <= atraccion.nextTurn+parseInt(atraccion.capacity); i++) {
      elements.push(<div className="grid-item">{i}</div>)
    }

    return (
        <div id="atraccion-container">
            <div id="turnos">
                <h1 id="atraccion-name">{atraccion.name}</h1>
                <div id="info-container">
                  <i className="fa-solid fa-light fa-people-line icon"></i>
                  <h2 id="title">Proximos Turnos</h2>
                </div>
                <div id="queue-container">
                  {elements}
                </div>
            </div>
            <div id="take-turn">
              <p id="take-text">Realiza tu registro acercando el código qr de tu manilla al sensor ubicado junto a esta pantalla</p>
              <form>
                <button type="submit" id="submit">Registro Turno</button>
              </form>
            </div>
        </div>
    )
  }
  