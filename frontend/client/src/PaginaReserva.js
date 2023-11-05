import React, { useState } from "react";
import Navbar from "./components/Navbar";
import DatePicker from "./components/DatePicker";
import "./components/Hero";
import "./DatePicker.css";
import { useNavigate } from "react-router-dom";

export default function PaginaReserva() {
  const [ciudad, setCiudad] = useState(""); // Estado para almacenar la ciudad
  const [fechas, setFechas] = useState({ inicio: null, fin: null }); // Estado para almacenar las fechas seleccionadas
  const navigate = useNavigate();

  const navbarLinks = [
    { url: "/", title: "Reservar" },
    { url: "/misreservas", title: "Mis Reservas" },
    { url: "/login", title: "Login" },
  ];

  const handleCiudadChange = (e) => {
    setCiudad(e.target.value);
  };

  const handleEnviarClick = () => {
    // Aquí puedes usar 'ciudad' y 'fechas' según tus necesidades
    // Por ejemplo, podrías enviar esta información a un servidor o realizar alguna otra acción.
    console.log("Ciudad:", ciudad);
    console.log("Fechas:", fechas);
  };

  return (
    <div className="PaginaReserva">
      <Navbar navbarLinks={navbarLinks} />
      <div className="date-container">
        <h1>Seleccione las fechas</h1>
        <DatePicker className="DatePicker" onSelect={(inicio, fin) => setFechas({ inicio, fin })} />
        <input
          type="text"
          placeholder="Ingrese la ciudad"
          value={ciudad}
          onChange={handleCiudadChange}
        />
        <button onClick={handleEnviarClick}>Enviar</button>
      </div>
    </div>
  );
}
