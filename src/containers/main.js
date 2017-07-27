import React from "react";
import { Tabs } from "antd";
import Demo01 from "./demo01";

export default class Containers extends React.Component {
  render() {
    let width = 512;
    let height = 512;
    const TabPane = Tabs.TabPane;
    return (
      <Tabs tabPosition="left" defaultActiveKey="1">
        <TabPane tab="1" key="1">
          <Demo01 width={width} height={height} />
        </TabPane>
      </Tabs>
    );
  }
}
