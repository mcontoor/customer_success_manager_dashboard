import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

import './Organisations.css'

function renderTable(headers, rows, nav) {
    const headerList = headers.map((val, _) => <td>{val}</td>)
    const rowsList = rows?.map(row => {
    return (
    <tr onClick={() => nav(`/organisation?orgName=${row.name}&id=${row.id}`)}>
        <td>{row.id}</td>
        <td>{row.name}</td>
        <td>{row.address}</td>
        <td>{row.created_at}</td>
        <td>${row.deal_amount}</td>
        <td>{row.days_till_renewal}</td>
    </tr>)})
    return (
        <table className='orgs-table' border={1}>
            <tr>
                {headerList}
            </tr>
            {rowsList}
        </table>
    )
}
const headers = ["Id", "Name", "Address", "Created at", "Deal Amount", "Days till renewal"]

export default function Organisations() {
    const [orgsList, setOrgsList] = useState([]);
    const [searchText, setSearchText] = useState("")
    const navigate = useNavigate();

    useEffect(() => {
        async function fetchOrgsData() {
            let url = 'http://localhost:8080/v1/organizations/'
            if (searchText.length) {
                url += `?q=${searchText}`
            }
            fetch(url)
        .then(response => response.json())
        .then(data => setOrgsList(data))
        .catch(e => console.log(e));
        }
        fetchOrgsData();
    }, [searchText])

    return (
        <div>
        <h1>Organisations</h1>
        <label>Search by Name: </label>
        <input 
            className='search-input'
            value={searchText} 
            onChange={(e) => setSearchText(e.target.value)} 
        />
        <div className='table-container'>
        {renderTable(headers, orgsList, navigate)}
        </div>
        </div>
    )
}