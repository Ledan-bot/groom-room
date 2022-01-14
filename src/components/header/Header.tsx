import React from 'react'
import { IconContext } from 'react-icons/'
import {FaPaw} from 'react-icons/fa'
import './Header.css'

const  Header: React.FC = () => {
  return (
    <>
      <header className="header-container">
        <div className="title-container">
          <h1 className="header-title">The Groom Room</h1>
          <IconContext.Provider value={{className: "header-icon"}}>
            <FaPaw />
          </IconContext.Provider>
        </div>
        <button className="header-btn">Schedule Appointment</button>
      </header>
    </>
  )
}
 export default Header