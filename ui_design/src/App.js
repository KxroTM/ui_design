import React, { useEffect, useState, useRef } from "react";
import "./App.css";

function App() {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState(0);
  
  const productsPerPage = 12;

  const loadingRef = useRef(false);

  useEffect(() => {
    //si requête en cours, ne pas envoyer de nouvelle requête
    if (loadingRef.current) return;

    loadingRef.current = true;
    setLoading(true);

    fetch(`http://localhost:8080/products?id_start=${page}&limit=${productsPerPage}`)
      .then((response) => response.json())
      .then((data) => {
        setProducts((prevProducts) => [...prevProducts, ...data.products]);
        setLoading(false);
        loadingRef.current = false;
      })
      .catch((error) => {
        console.error("Erreur lors de la récupération des produits :", error);
        setLoading(false);
        loadingRef.current = false;
      });
  }, [page]);

  //si l'utilisateur atteint le bas de la page
  const handleScroll = () => {
    const bottom = window.innerHeight + document.documentElement.scrollTop === document.documentElement.offsetHeight;
    if (bottom && !loadingRef.current) {
      
      setPage((prevPage) => prevPage + 1);
    }
  };

  useEffect(() => {
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  if (loading && page === 1) {
    return <p>Chargement des produits...</p>;
  }

  return (
    <div>
      <h1 class="title">Produits</h1>
      <div class="all-products">
        {products.map((product) => (
          <div class="product" key={product.id}>  
              <h2 class="product-title">{product.name}</h2>
              <img src={product.image_url} alt={product.name} width="100" />
              <p class="product-description">{product.description}</p>
              <p class="price"> ${parseFloat(product.price).toFixed(2)}</p>
          </div>
        ))}
      </div>

      {loading && <p>Chargement...</p>}
    </div>
  );
}

export default App;
