import React, { useEffect, useState } from 'react';
import { useSearchParams } from 'react-router-dom';

export default function Organisation() {
    const [searchParams] = useSearchParams();
    const id = searchParams.get('id')

    const [orgData, setOrgData] = useState({});
    const [usersList, setUsersList] = useState([]);
    const [loadState, setLoadState] = useState("initial")

    useEffect(() => {
        async function fetchOrgData() {
        setLoadState("loading")
        fetch(`http://localhost:8080/v1/organizations/${id}/org-data`)
        .then(response => response.json())
        .then(data => setOrgData(data))
        .then(() => setLoadState("complete"))
        .catch(e => console.log(e));
        }
        fetchOrgData();
    }, [])

    useEffect(() => {
        async function fetchUsersList() {
        fetch(`http://localhost:8080/v1/organizations/${id}/users`)
        .then(response => response.json())
        .then(data => setUsersList(data))
        .catch(e => console.log(e));
        }
        fetchUsersList();
    }, [])



    return (
        <div>
            
        { loadState === "loading" ? <h1>loading....</h1> : (
        <>
            <h1>{orgData.name}</h1>
            <div>{JSON.stringify(orgData)}</div>
            <hr />
            <h1>users list</h1>
            <div>{JSON.stringify(usersList)}</div>
        </>
)}

        </div>
    )
}
