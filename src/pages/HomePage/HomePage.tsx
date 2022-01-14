import React from 'react'
import Header from '../../components/header/Header'
import './HomePage.css'

const HomePage: React.FC = () => {
  return (
    <>
      <Header />
      <main className="main-container">
        <h1 className="main-title">Your Neighborhood Pet Groomers - Handling All of Your Grooming Needs!</h1>
        <h2 className="services-title">Services</h2>
        <section className="sub-main-container">
          <img src="/assets/finalcut.jpg" alt="dog image"></img>
          <h3>Full Service Bathing 	&#38; Grooming For All Your Animal Needs</h3>
        </section>
        <section className="sub-main-container">
          <img src="/assets/fuzz.jpg"></img>
          <h3>Weekly &#38; Touch-ups to Keep Your Family Members Looking Their Best!</h3>
        </section>
      </main>
    </>
  )
}

export default HomePage