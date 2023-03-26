import { Layout } from "antd";
import React from "react";
import Table from "./table";
const { Header, Footer, Sider, Content } = Layout;

const App = () => (
  <Layout>
    {/* <Header
      style={{
        height: "10vh",
      }}
    >
      Header
    </Header> */}
    <Content
      style={{
          margin: "auto",
        // height: "80vh",
        width: "80vw",
        // overflow: "scroll",
      }}
    >
      <Table />
    </Content>
    {/* <Footer
      style={{
        height: "10vh",
      }}
    >
      Footer
    </Footer> */}
  </Layout>
);

export default App;
