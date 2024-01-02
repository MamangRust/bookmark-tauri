import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

export default function AllFolder() {
  const [folders, setFolders] = useState([]);

  useEffect(() => {
    axios
      .get('http://localhost:8000/folder') // Ganti URL sesuai dengan endpoint yang benar
      .then((response) => {
        setFolders(response.data);
      })
      .catch((error) => {
        console.error('Error fetching folders:', error);
      });
  }, []);

  return (
    <div className="container mt-4">
      <div className="row">
        {folders.map((folder, index) => (
          <div className="col-md-4 mb-4" key={index}>
            <div className="card">
              <img
                src="https://via.placeholder.com/150"
                className="card-img-top"
                alt="Folder"
              />
              <div className="card-body">
                <h5 className="card-title">{folder.name_folder}</h5>

                <p className="card-text">Folder Description or Content</p>
                <Link
                  className="btn btn-primary"
                  to={'/all-text/' + folder.name_folder}
                >
                  Open Folder
                </Link>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
