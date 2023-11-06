import React, { useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom";

const DisponibilidadnReserva = () => {
  const [formData, setFormData] = useState({
    name: "",
    roomAmount: 0,
    description: "",
    city: "",
    streetName: "",
    streetNumber: 0,
    rate: 0.0,
    amenities: {
      // Define las propiedades de Amenities aquí según la estructura de tu base de datos
    },
    images: [],
  });

  const [deleteId, setDeleteId] = useState("");
  const [updateId, setUpdateId] = useState("");

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post("http://localhost:8080/hotel", formData);
      alert("Hotel insertado con éxito");
      // Limpia el formulario después de la inserción
      setFormData({
        name: "",
        roomAmount: 0,
        description: "",
        city: "",
        streetName: "",
        streetNumber: 0,
        rate: 0.0,
        amenities: {},
        images: [],
      });
    } catch (error) {
      console.error("Error al insertar el hotel:", error);
    }
  };

  const handleDelete = async () => {
    try {
      await axios.delete(`http://localhost:8080/hotel/${deleteId}`);
      alert("Hotel eliminado con éxito");
    } catch (error) {
      console.error("Error al eliminar el hotel:", error);
    }
  };

  const handleUpdate = async () => {
    try {
      await axios.put(`http://localhost:8080/hotel${updateId}`, formData);
      alert("Hotel actualizado con éxito");
    } catch (error) {
      console.error("Error al actualizar el hotel:", error);
    }
  };

  return (
    <div className="confirmacion-reserva">
      <h2>Insertar Nuevo Hotel</h2>
      <form onSubmit={handleSubmit}>
        <label htmlFor="name">Nombre:</label>
        <input type="text" id="name" name="name" onChange={handleInputChange} value={formData.name} required />

        <label htmlFor="roomAmount">Número de Habitaciones:</label>
        <input type="number" id="roomAmount" name="roomAmount" onChange={handleInputChange} value={formData.roomAmount} required />

        <label htmlFor="description">Descripción:</label>
        <input type="text" id="description" name="description" onChange={handleInputChange} value={formData.description} required />

        <label htmlFor="city">Ciudad:</label>
        <input type="text" id="city" name="city" onChange={handleInputChange} value={formData.city} required />

        <label htmlFor="streetName">Nombre de la Calle:</label>
        <input type="text" id="streetName" name="streetName" onChange={handleInputChange} value={formData.streetName} required />

        <label htmlFor="streetNumber">Número de la Calle:</label>
        <input type="number" id="streetNumber" name="streetNumber" onChange={handleInputChange} value={formData.streetNumber} required />

        <label htmlFor="rate">Tarifa:</label>
        <input type="number" id="rate" name="rate" step="0.01" onChange={handleInputChange} value={formData.rate} required />

        {/* Agrega campos adicionales según las propiedades de Hotel */}
        {/* ... */}

        <button type="submit">Insertar Hotel</button>
      </form>

      <h2>Eliminar Hotel</h2>
      <label htmlFor="deleteId">ID del hotel a eliminar:</label>
      <input type="text" id="deleteId" onChange={(e) => setDeleteId(e.target.value)} />
      <button onClick={handleDelete}>Eliminar Hotel</button>

      {/* Agrega campos adicionales para la actualización */}
      {/* ... */}
    
    <h2>Actualizar Hotel</h2>
    <label htmlFor="updateId">ID del hotel a actualizar:</label>
    <input type="text" id="updateId" onChange={(e) => setUpdateId(e.target.value)} />
    
    <label htmlFor="updateName">Nuevo Nombre:</label>
    <input type="text" id="updateName" name="name" onChange={handleInputChange} value={formData.name} required />
    
    <label htmlFor="updateRoomAmount">Nuevo Número de Habitaciones:</label>
    <input type="number" id="updateRoomAmount" name="roomAmount" onChange={handleInputChange} value={formData.roomAmount} required />
    
    <label htmlFor="updateDescription">Nueva Descripción:</label>
    <input type="text" id="updateDescription" name="description" onChange={handleInputChange} value={formData.description} required />
    
    <label htmlFor="updateCity">Nueva Ciudad:</label>
    <input type="text" id="updateCity" name="city" onChange={handleInputChange} value={formData.city} required />
    
    <label htmlFor="updateStreetName">Nuevo Nombre de Calle:</label>
    <input type="text" id="updateStreetName" name="streetName" onChange={handleInputChange} value={formData.streetName} required />
    
    <label htmlFor="updateStreetNumber">Nuevo Número de Calle:</label>
    <input type="number" id="updateStreetNumber" name="streetNumber" onChange={handleInputChange} value={formData.streetNumber} required />
    
    <label htmlFor="updateRate">Nueva Tarifa:</label>
    <input type="number" id="updateRate" name="rate" step="0.01" onChange={handleInputChange} value={formData.rate} required />
    
    {/* Agrega campos adicionales según las propiedades de Hotel */}
    {/* ... */}
    
    <button onClick={handleUpdate}>Actualizar Hotel</button>
    
    
    
    <Link to="/">
              <button>Volver al menu principal</button>
            </Link>
    </div>

    
  );
};


export default DisponibilidadnReserva;
