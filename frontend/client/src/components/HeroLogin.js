import React from "react";
import "./HeroLogin.css";

const HeroLogin = ({ imageSrc }) => {
  return (
    <div className="hero">
      <img src={imageSrc} alt=" MyStay" className="hero__image" />
    </div>
  );
};

export default HeroLogin;