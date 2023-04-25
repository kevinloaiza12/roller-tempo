import Image from 'next/image'
import Link from 'next/link'
import placeholder from '../../../public/placeholder.jpg'
import './atraccion.css'
export const metadata = {
    title: 'Usuarios',
  }

// const atracciones = [
//   {
//     "capacity": "50",
//     "description": "Fire",
//     "duration": "60",
//     "name": "atraccion",
//     "nextTurn": 50,
//   },
//   {
//     "capacity": "20",
//     "description": "water",
//     "duration": "6",
//     "name": "atraccion extrema",
//     "nextTurn": 501,

//   },
//   {
//     "capacity": "1",
//     "description": "Earth",
//     "duration": "1",
//     "name": "atraccion zzzzz",
//     "nextTurn": 10,
//   }
// ]

async function getData(param) {
  let url = "http://127.0.0.1:3000/api/attractioninfo/" + param.params.atraccion;

  let response = await
  fetch(url)
  let data = await response.json()
  return data;
  }

export default async function Atraccion(param) {
    //const atraccion = atracciones[(param.params.atraccion)] 
    const atraccion = await getData(param);
    return (
      <main>
        <header>
            <Link href="/atraccion">
                <i className="fa-solid fa-arrow-left fa-xl header-back"></i>
            </Link>
            <h1 id="header-text"> Atraccion </h1>
        </header>
        <div id="atraccion-container">
            <Image id="atraccion-image" src={placeholder}/>
            <h1 id="atraccion-title">{atraccion.name}</h1>
            <p id="atraccion-desc">{atraccion.description}</p>
            <h1 id="atraccion-title">Fila</h1>
            <div id="fila-container">
              <div className="fila-sub-container">
                <i class="fa-regular fa-clock icon"></i>
                <p className="icon-text">{atraccion.duration} min</p>
              </div>
              <div className="fila-sub-container">
                <i className="fa-solid fa-light fa-people-line icon"></i>
                <p className="icon-text">Turnos: {atraccion.nextTurn} - {atraccion.nextTurn+atraccion.capacity}</p>
                </div>
            </div>
          <form id="user-turn" action="/usuarios" method="get">
            <input type="submit" value="Pedir turno" id="user-turn-button"></input>
          </form>
        </div>
      </main>
    )
  }
  