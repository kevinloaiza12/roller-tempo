import Image from 'next/image'
import Link from 'next/link'
import './premios.css'
import placeholder from '../../public/premio_placeholder.jpg';
export const metadata = {
    title: 'Usuarios',
  }

// const premios = [
//   {
//     "price": "50",
//     "description": "Fire",
//     "name": "atraccion",
//   },
//   {
//     "description": "water",
//     "price": "6",
//     "name": "atraccion extrema",

//   },
//   {
//     "description": "Earth",
//     "price": "1",
//     "name": "atraccion zzzzz",
//   }
// ]

async function getData() {
  let response = await
  fetch("http://127.0.0.1:3000/rewards")
  let data = await response.json()
  return data;
  }

export default async function Premios() {
  const premios = await getData();
  const listPremios = premios.map((premio, index) =>{
    const route = "/premios/" + premio.name
    return(
      <Link href={route}>
        <div className="premios-container">
          <Image className="premios-image"src={placeholder}/>
          <div className="premios-info">
            <h1 className="premios-title">{premio.name}</h1>
            <div id="price_container">
              <i class="fa-sharp fa-solid fa-coins price_icon_2"></i>
              <h1 className="premios_title_2">{premio.price}</h1>
              
            </div>

          </div>
        </div>
      </Link>
    )
  })
    return (
      <main id="premios">
        <header>
          <h1 id="header-text"> Premios </h1>
        </header>
        {listPremios}
      </main>
    )
  }