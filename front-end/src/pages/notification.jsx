import * as React from 'react';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import {useEffect} from "react";
import {setUser} from './login'
import Button from "@mui/material/Button";
import {useLocation} from "react-router-dom";
import { DataGrid } from '@mui/x-data-grid';

function createData(
    coordinate,
    distance,
    quality,
    type,
    time,
    frequency,
    state
) {
    return { coordinate,
        distance,
        quality,
        type,
        time,
        frequency,
        state };
}



const columns = [
    { field: 'id', headerName: 'ID', width: 70 },
    {
        field: 'coordinate',
        headerName: 'coordinate',

        width: 190,
    },
    {
        field: 'quality',
        headerName: 'quality',
        type: 'number',
        width: 190,
    },
    {
        field: 'type',
        headerName: 'type',

        width: 190,
    },
    {
        field: 'time',
        headerName: 'time',
        type: 'number',
        width: 190,
    },
    {
        field: 'frequency',
        headerName: 'frequency',
        type: 'number',
        width: 190,
    },
    {
        field: 'state',
        headerName: 'state',
        type: 'number',
        width: 190,
    },
    // {
    //     field: 'fullName',
    //     headerName: 'Full name',
    //     description: 'This column has a value getter and is not sortable.',
    //     sortable: false,
    //     width: 160,
    //     valueGetter: (params) =>
    //         `${params.row.firstName || ''} ${params.row.lastName || ''}`,
    // },
];

const rows = [
    {
        id: 2,
        coordinate: "asdf11111adf323",
        distance: 412.7,
        quality: 0.888,
        type: "sunset",
        time: 1673845785,
        frequency: 2,
        state: 1,
    },
    // { id: 2, lastName: 'Lannister', firstName: 'Cersei', age: 42 },
    // { id: 3, lastName: 'Lannister', firstName: 'Jaime', age: 45 },
    // { id: 4, lastName: 'Stark', firstName: 'Arya', age: 16 },
    // { id: 5, lastName: 'Targaryen', firstName: 'Daenerys', age: null },
    // { id: 6, lastName: 'Melisandre', firstName: null, age: 150 },
    // { id: 7, lastName: 'Clifford', firstName: 'Ferrara', age: 44 },
    // { id: 8, lastName: 'Frances', firstName: 'Rossini', age: 36 },
    // { id: 9, lastName: 'Roxie', firstName: 'Harvey', age: 65 },
];

export default function CurrentNotification() {

    let location = useLocation()
    useEffect(()=>{

            const user_id = location.state.name
            const user_autoken = location.state.tok.user_token.token
            fetch("http://ec2-3-83-223-34.compute-1.amazonaws.com:8800/api/v1/rule",{
                method:"POST",
                headers:{Authorization:user_autoken},
                body:JSON.stringify(user_id)

            }).then(res=>res.json())
                .then((res)=>{
                    if (res){
                        let result = []
                        result =res
                        alert("noti successfully");
                       }

                })
        })

    return (

        <div style={{ height: 400, width: '100%' }}>
            <DataGrid
                rows={rows}
                columns={columns}
                pageSize={5}
                rowsPerPageOptions={[5]}
                checkboxSelection
            />
            <Button
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
                // onClick={handleCheckCurRule}
            >
                Set new rules
            </Button>
            <Button
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
                // onClick={handleCheckCurRule}
            >
                update rules
            </Button>
            <Button
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
                // onClick={handleCheckCurRule}
            >
                delete rules
            </Button>
        </div>



    );
}