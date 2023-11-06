import React from "react";
import { Link } from "react-router-dom";

const navbarLinks = [
  { url: "/", title: "Reservar" },
  { url: "/misreservas", title: "Mis Reservas" },
  { url: "/login", title: "Login" },
];

const DisponibilidadnReserva = () => {
  return (
    <div className="confirmacion-reserva">
      <h2>Informacion de su hotel</h2>
      <Link to="/">
        <button>Volver al menú principal</button>
      </Link>
    </div>
  );
};

export default DisponibilidadnReserva;