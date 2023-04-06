import React, { useEffect, useState } from 'react';
import { useSearchParams } from 'react-router-dom';
import ReactECharts from 'echarts-for-react';
import './Organisation.css';

function renderTable(headers, rows = []) {
    const headerList = headers.map((val, _) => <td>{val}</td>)
    const rowsList = rows?.map(row => {
    return (
    <tr>
        <td>{row.first_name + " " + row.last_name}</td>
        <td>{row.email}</td>
        <td>{row.gender}</td>
        <td>{row.age}</td>
        <td>{row.active_hours_on_app}</td>
        <td>{row.has_unsubscribed}</td>
    </tr>)})
    return (
        <table border={1} className="org-table">
            <tr>
                {headerList}
            </tr>
            {rowsList}
        </table>
    )
}

const headers = ["Name", "Email", "Gender", "Age", "Active Hours On App", "Has unsubscribed"]

export default function Organisation() {
    const [searchParams] = useSearchParams();
    const id = searchParams.get('id')

    const [orgData, setOrgData] = useState({});
    const [usersList, setUsersList] = useState([]);
    const [loadState, setLoadState] = useState("initial")
    const [hrsOnAppHist, setHrsOnAppHist] = useState({});
    const [sortBy, setSortBy] = useState("");
    const [order, setOrder] = useState("ASC");
    const [page, setPage] = useState(1);

    useEffect(() => {
        async function fetchOrgData() {
        setLoadState("loading")
        fetch(`http://localhost:8080/v1/organizations/${id}/org-data`)
        .then(response => response.json())
        .then(data => {
            setOrgData(data)
            const keysArr = [];
            const valsArr = [];
            data?.hours_on_product.forEach(pair => {
                keysArr.push(pair.hourrange + " hrs")
                valsArr.push(pair.count)
            });
            setHrsOnAppHist({ keys: keysArr, vals: valsArr})
        })
        .then(() => setLoadState("complete"))
        .catch(e => console.log(e));
        }
        fetchOrgData();
    }, [])

    useEffect(() => {
        async function fetchUsersList() {
            let url = `http://localhost:8080/v1/organizations/${id}/users?page=${page}`;
            if (sortBy.length) {
                url += `&sortBy=${sortBy}&order=${order}`
            }
        fetch(url)
        .then(response => response.json())
        .then(data => setUsersList(data))
        .catch(e => console.log(e));
        }
        fetchUsersList();
    }, [page, sortBy, order])

    const option = {
        title: {
          text: 'Hours spent on product'
        },
        tooltip: {},
        legend: {
          data:['what i want it to be']
        },
        xAxis: {
          data: hrsOnAppHist?.keys || []
        },
        yAxis: {
        },
        series: [{
          name: 'Hour Range',
          type: 'bar',
          data: hrsOnAppHist?.vals || []
        }]
      };

      console.log(sortBy, order)

      function onChangeColOption(e) {
            let x = e.target.value;
            if (x === "Name") {
                setSortBy("first_name")
            } else {
                setSortBy(x.split(' ').join('_').toLowerCase())
            }
      }

    
    return (
        <div>
        { loadState === "loading" ? <h1>loading....</h1> : (
        <>
            <h1>{orgData.name}</h1>
            <div className='organisation-info'>
            <div>
                <h2>Address</h2>
                <br />
                <p>{orgData.address}</p>
            </div>
            <div>
                <h2>Deal Amount</h2>
                <br />
                <p>{orgData.deal_amount}</p>
            </div>
            <div>
                <h2>Deal Amount</h2>
                <br />
                <p>{orgData.days_till_renewal}</p>
                </div>
            </div>
            <ReactECharts
                option={option}
                style={{ height: 250 }}
                opts={{ renderer: 'svg' }}
            />
            <hr />
            <h1 className='users-list'>Users list</h1>
            <div className='pagination'>
                <label>Sort:</label>
                <select onChange={onChangeColOption}>
                    <option value="">Select an option</option>
                    {headers.map(header => <option value={header}>{header}</option>)}
                </select>
                <label>Order:</label>
                <select onChange={(e) => setOrder(e.target.value)}>
                    <option value="ASC">Ascending</option>
                    <option value="DESC">Descending</option>
                </select>
            </div>
            {renderTable(headers, usersList)}
            <div className='pagination'>
            <button disabled={page === 1} onClick={() => setPage(prev => prev - 1)}>Prev Page</button>
            <p>{page} / {usersList.pageCount || 5}</p>
            <button onClick={() => setPage(prev => prev + 1)}>Next Page</button>
            </div>
        </>
        )}
        </div>
    )
}
