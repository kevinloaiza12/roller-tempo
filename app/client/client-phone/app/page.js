import { Inter } from 'next/font/google'
import logo from '../public/logo.png'
import Image from 'next/image'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <main id="welcome_page">
      <Image id="logo"src={logo}></Image>
      <h1 id="logo_title">Roller Tempo</h1> 
    </main>
  )
}
