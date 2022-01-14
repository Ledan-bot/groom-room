import React from 'react'
import Header from '../../components/header/Header'
import Footer from '../../components/footer/Footer'
import './HomePage.css'

const HomePage: React.FC = () => {
  return (
    <>
      <Header />
      <main className="main-container">
        <h1 className="main-title">Your Neighborhood Pet Groomers - Handling All of Your Grooming Needs!</h1>
        <h2 className="services-title">Services</h2>
        <section className="sub-main-container">
          <img src="/assets/finalcut.jpg" alt="dog image" className="img-homepage" ></img>
          <h3>Full Service Bathing	&#38; Grooming For All Your Animal Needs</h3>
        </section>
        <section className="sub-main-container-reverse">
          <img src="/assets/fuzz.jpg" alt="dog image" className="img-homepage" ></img>
          <h3>Weekly &#38; Monthly Touch-ups to Keep Your Family Members Looking Their Best!</h3>
        </section>
        <section className="sub-main-container">
          <img src="/assets/headlights.jpg" alt="dog image" className="img-homepage"></img>
          <h3>Grooming is Essential to Keeping your Pet Healthy!</h3>
        </section>
      </main>
      <Footer />
    </>
  )
}

export default HomePage