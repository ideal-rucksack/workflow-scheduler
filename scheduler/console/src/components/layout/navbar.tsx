import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';
import {Actions} from "@/components/layout/index";
import {logo} from "@/assets";

interface NavBarProps {
  $isScrolled: boolean;
}

const NavBar = styled.div<NavBarProps>`
  position: ${props => props.$isScrolled ? 'fixed' : 'absolute'};
  top: 0;
  width: 100%;
  display: flex;
  justify-content: space-between;
  padding: 20px;
  background-color: #fff;
  box-shadow: ${props => props.$isScrolled ? '0px 5px 10px rgba(0, 0, 0, 0.15)' : 'none'};
  transition: all 0.5s ease-in-out;
  z-index: 1000;
  filter: ${props => props.$isScrolled ? 'drop-shadow(0px 0px 5px rgba(0, 0, 0, 0.15))' : 'none'};
`;

const Logo = styled.img`
  height: 40px;
`;

const SearchInput = styled.input`
  margin-right: 20px;
`;

const Navbar = () => {
  const [isScrolled, setIsScrolled] = useState<boolean>(false);

  useEffect(() => {
    const handleScroll = () => {
      const isScrolled = window.scrollY > 100;
      setIsScrolled(isScrolled);
    };

    document.addEventListener('scroll', handleScroll);
    return () => {
      document.removeEventListener('scroll', handleScroll);
    };
  }, []);

  return (
    <NavBar $isScrolled={isScrolled}>
      <div>
        <Link to="/">
          <Logo src={logo} alt="Logo" />
        </Link>
      </div>
      <div style={{ display: 'flex', alignItems: 'center' }}>
        <Actions />
      </div>
    </NavBar>
  );
};

export default Navbar;
