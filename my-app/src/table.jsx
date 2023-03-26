import { Table, Space, Tag, Button, message } from "antd";
import qs from "qs";
import React, { useEffect, useState } from "react";
import moment from "moment";

const copyTextToClipboard = async (text) => {
  if ("clipboard" in navigator) {
    return await navigator.clipboard.writeText(text);
  } else {
    return document.execCommand("copy", true, text);
  }
};

const Tables = () => {
  const columns = [
    { title: "ID", dataIndex: "ID", ellipsis: true, width: 50 },
    {
      title: "Text",
      dataIndex: "Text",
      ellipsis: true,
      render: (_, record) => {
        return (
          <div  onClick={(e) => copyTextToClipboard(record.Text, e)}>
            {record.Text}
          </div>
        );
      },
    },
    {
      title: "CreatedAt",
      dataIndex: "CreatedAt",
      ellipsis: true,
      width: 150,
      render: (time) => time2date(time),
    },
    {
      title: "UpdatedAt",
      dataIndex: "UpdatedAt",
      ellipsis: true,
      width: 150,
      render: (time) => time2date(time),
    },
  ];

  const time2date = (time) => {
    return moment(time).format("YY-MM-DD HH:mm:ss");
  };
  const getRandomuserParams = (params) => ({
    PageSize: params.pagination?.pageSize,
    PageNum: params.pagination?.current,
    ...params,
  });

  const [data, setData] = useState();
  const [loading, setLoading] = useState(false);
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 100,
  });

  const fetchData = (params = {}) => {
    setLoading(true);
    fetch(`/api?${qs.stringify(getRandomuserParams(params))}`)
      .then((res) => res.json())
      .then((results) => {
        console.log("this is res", results);
        setData(results.data);
        setLoading(false);
        setPagination({
          ...params.pagination,
          total: results.count, // 200 is mock data, you should read it from server
          // total: data.totalCount,
          pageSizeOptions: [10, 50, 100, 500, 1000],
        });
      });
  };

  useEffect(() => {
    fetchData({
      pagination,
    });
  }, []);

  const handleTableChange = (newPagination, filters, sorter) => {
    fetchData({
      sortField: sorter.field,
      sortOrder: sorter.order,
      pagination: newPagination,
      ...filters,
    });
  };

  return (
    <Table
      columns={columns}
      rowKey={(record) => record.ID}
      dataSource={data}
      pagination={pagination}
      loading={loading}
      onChange={handleTableChange}
      size="small"
      scroll={{
        scrollToFirstRowOnChange: true,
        // x: "50vw",
        // y: "70vh"
      }}
    />
  );
};

export default Tables;
