import Image from 'next/image'
import Link from 'next/link'
import './atraccion.css'
import placeholder from '../../public/placeholder.jpg';
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

async function getData() {
  let response = await
  fetch("http://127.0.0.1:3000/attractions")
  let data = await response.json()
  return data;
  }

export default async function Atraccion() {
  const atracciones = await getData();
  const listAtracciones = atracciones.map((atraccion, index) =>{
    const route = "/atraccion/" + atraccion.name
    return(
      <Link href={route}>
        <div className="atraccion-container">
          <Image className="atraccion-image"src={placeholder}/>
          <div className="atraccion-info">
            <h1 className="atraccion-title">{atraccion.name}</h1>
          </div>
        </div>
      </Link>
    )
  })
    return (
      <main id="atracciones">
        <header>
          <h1 id="header-text"> Atracciones </h1>
        </header>
        {listAtracciones}
      </main>
    )
  }
  