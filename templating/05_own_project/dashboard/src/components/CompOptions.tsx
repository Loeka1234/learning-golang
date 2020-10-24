import React, { useState } from "react";
import styled from "styled-components";
import { ComponentOptions } from "../types/webconfig.types";
import axios from "axios";

const ColorWrapper = styled.div`
  display: flex;
  justify-content: flex-end;
  align-items: center;
  width: 85%;
  margin: 5px 0;
`;

const Label = styled.label`
  margin-right: 5px;
`;

const ColorInput = styled.input`
  border: none;
  outline: none;
  -webkit-appearance: none;
  cursor: pointer;
  width: 44px;
  height: 44px;
  &::-webkit-color-swatch-wrapper {
    padding: 0;
  }
  &::-webkit-color-swatch {
    border: none;
  }
`;

const Button = styled.button`
  padding: 5px 7px;
  border: none;
  color: #3a404a;
  background: white;
  margin-top: 4px;
  border-radius: 3px;
  cursor: pointer;
  transition: background 0.3s ease-in-out;
  &:hover {
    background: #c0c0c0;
  }
`;

export interface CompOptionsProps {
  options: ComponentOptions;
  componentType: string;
  reloadIFrame: () => void;
}

export const CompOptions: React.FC<CompOptionsProps> = ({
  options,
  componentType,
  reloadIFrame,
}) => {
  const [optionsInState, setOptionsInState] = useState(options);

  const handleSaveOptions = async () => {
    const res = await axios.post(
      `/editoptions?type=${componentType}`,
      optionsInState
    );
    if (res.status === 200) reloadIFrame();
  };

  return (
    <>
      <ColorWrapper>
        <Label htmlFor="backColor">Background Color: </Label>
        <ColorInput
          type="color"
          defaultValue={options.backColor}
          id="backColor"
          color={optionsInState.backColor}
          onChange={e =>
            setOptionsInState({
              ...options,
              backColor: e.target.value,
            })
          }
        />
      </ColorWrapper>
      <ColorWrapper>
        <Label htmlFor="color">Color: </Label>
        <ColorInput
          type="color"
          defaultValue={options.color}
          color={optionsInState.color}
          onChange={e =>
            setOptionsInState({
              ...options,
              color: e.target.value,
            })
          }
        />
      </ColorWrapper>
      <Button onClick={handleSaveOptions}>Save Options</Button>
    </>
  );
};
