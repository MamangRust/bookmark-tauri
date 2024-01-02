import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link, useParams } from 'react-router-dom';

export default function AllMarkdown() {
  const { folder } = useParams();
  const [markdowns, setMarkdowns] = useState([]);

  useEffect(() => {
    axios
      .get(`http://localhost:8000/folder/find/${folder}`)
      .then((response) => {
        setMarkdowns(response.data);
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  return (
    <div>
      <h1>All Markdown</h1>
      <hr />
      <div className="card-container">
        {markdowns.map((markdown, index) => (
          <div className="card" key={index}>
            <div className="card-body">
              <h2 className="card-title">{markdown.title}</h2>
              <p className="card-text">{markdown.content}</p>

              <Link
                to={'/find-file/' + folder + '/' + markdown.title}
                className="btn btn-primary"
              >
                View Details
              </Link>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
