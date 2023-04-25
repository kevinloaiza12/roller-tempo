import Image from 'next/image'
import Link from 'next/link'
import placeholder from '../../../public/premio_placeholder.jpg'
import './premio.css'
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


async function getData(param) {
  let url = "http://127.0.0.1:3000/api/rewardinfo/" + param.params.premio;
  let response = await
  fetch(url)
  let data = await response.json()
  return data;
  }

export default async function Premio(param) {
    //const premio = premios[(param.params.premio)] 
    const premio = await getData(param);
    return (
      <main>
        <header>
            <Link href="/premios">
                <i className="fa-solid fa-arrow-left fa-xl header-back"></i>
            </Link>
            <h1 id="header-text"> Premios </h1>
        </header>
        <div id="premio-container">
            <Image id="premio-img" src={placeholder}/>
            <h1 id="premio-title">{premio.name}</h1>
            <p id="premio-desc">{premio.description}</p>
            <div id="price">
              <i class="fa-sharp fa-solid fa-coins price_icon"></i>
              <p id="price_number"> {premio.price}</p>
            </div>
        </div>
      </main>
    )
  }
  