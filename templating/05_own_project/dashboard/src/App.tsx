import React, { useEffect, useRef, useState } from "react";
import styled from "styled-components";
import axios from "axios";
import { Webconfig } from "./types/webconfig.types";
import { CompOptions } from "./components/CompOptions";

import "./global.css";

const Main = styled.main`
  width: 100%;
  height: 100vh;
`;

const Settings = styled.div`
  width: 217px;
  height: 100%;
  float: left;
  background: #3a404a;
  color: white;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  flex-direction: column;
  border-right: 3px solid #32363d;
`;

const Header = styled.h1`
  margin-top: 15px;
  font-weight: 500;
`;

const Components = styled.div`
  width: 100%;
`;

const ComponentWrapper = styled.div`
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  margin-top: 10px;
`;

const ComponentSelection = styled.select`
  cursor: pointer;
  border: 2px solid white;
  border-radius: 3px;
  color: white;
  background: #3a404a;
  padding: 3px 6px;
  margin: 10px 0 0 0;
  &:focus {
    outline: none;
  }
  option:hover {
    background: #586170;
  }
`;

const Wrapper = styled.div`
  width: calc(100% - 220px);
  height: 100%;
  float: left;
  background: grey;
  iframe {
    height: 100%;
    width: 100%;
    border: none;
  }
`;

const App = () => {
  const [headers, setHeader] = useState<null | string[]>(null);
  const [sections, setSections] = useState<null | string[]>(null);
  const [footers, setFooters] = useState<null | string[]>(null);
  const [webconfig, setWebconfig] = useState<null | Webconfig>(null);
  const iframeRef = useRef<HTMLIFrameElement>(null);

  useEffect(() => {
    const getAll = async () => {
      const headers = await axios.get("/headers");
      const sections = await axios.get("/sections");
      const footers = await axios.get("/footers");
      const webconfig = await axios.get("/webconfig");
      setHeader(headers.data);
      setSections(sections.data);
      setFooters(footers.data);
      setWebconfig(webconfig.data);
    };

    getAll();
  }, []);

  const handleChange = async (
    e: React.ChangeEvent<HTMLSelectElement>,
    componentType: string
  ) => {
    await axios.post(
      `/edit?type=${componentType}&comp=${e.target.value}.gohtml`
    );

    reloadIFrame();
  };

  const reloadIFrame = () => {
    if (iframeRef.current) {
      iframeRef.current.src = "http://localhost:3001";
    }
  };

  return headers && sections && footers && webconfig ? (
    <Main>
      <Settings>
        <Header>Options</Header>
        <Components>
          <ComponentWrapper>
            <ComponentSelection
              onChange={e => handleChange(e, "header")}
              defaultValue={webconfig.header.selected.replace(".gohtml", "")}
            >
              {headers.map(name => (
                <option key={name}>{name.replace(".gohtml", "")}</option>
              ))}
            </ComponentSelection>
            <CompOptions
              options={webconfig.header.options}
              componentType="header"
              reloadIFrame={reloadIFrame}
            />
          </ComponentWrapper>

          <ComponentWrapper>
            <ComponentSelection
              onChange={e => handleChange(e, "section")}
              defaultValue={webconfig.section.selected.replace(".gohtml", "")}
            >
              {sections.map(name => (
                <option key={name}>{name.replace(".gohtml", "")}</option>
              ))}
            </ComponentSelection>
            <CompOptions
              options={webconfig.section.options}
              componentType="section"
              reloadIFrame={reloadIFrame}
            />
          </ComponentWrapper>

          <ComponentWrapper>
            <ComponentSelection
              onChange={e => handleChange(e, "footer")}
              defaultValue={webconfig.footer.selected.replace(".gohtml", "")}
            >
              {footers.map(name => (
                <option key={name}>{name.replace(".gohtml", "")}</option>
              ))}
            </ComponentSelection>
            <CompOptions
              options={webconfig.footer.options}
              componentType="footer"
              reloadIFrame={reloadIFrame}
            />
          </ComponentWrapper>
        </Components>
      </Settings>
      <Wrapper>
        <iframe
          src="http://localhost:3001"
          ref={iframeRef}
          title="Web page"
        ></iframe>
      </Wrapper>
    </Main>
  ) : null;
};

export default App;
