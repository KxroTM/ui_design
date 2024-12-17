import React, { useEffect, useState, useRef } from "react";

function App() {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState(0);
  
  const productsPerPage = 10;

  // Utiliser useRef pour suivre si la requête est en cours
  const loadingRef = useRef(false);

  useEffect(() => {
    // Si une requête est déjà en cours, ne pas envoyer de nouvelle requête
    if (loadingRef.current) return;

    // Marquer que le chargement commence
    loadingRef.current = true;
    setLoading(true);

    fetch(`http://localhost:8080/products?id_start=${page}&limit=${productsPerPage}`)
      .then((response) => response.json())
      .then((data) => {
        setProducts((prevProducts) => [...prevProducts, ...data.products]);
        setLoading(false);  // Fin du chargement
        loadingRef.current = false; // Réinitialiser le flag de chargement
      })
      .catch((error) => {
        console.error("Erreur lors de la récupération des produits :", error);
        setLoading(false);  // Fin du chargement même en cas d'erreur
        loadingRef.current = false;
      });
  }, [page]);  // Dépend de la page uniquement

  // Détecter si l'utilisateur atteint le bas de la page
  const handleScroll = () => {
    const bottom = window.innerHeight + document.documentElement.scrollTop === document.documentElement.offsetHeight;
    if (bottom && !loadingRef.current) {
      setPage((prevPage) => prevPage + 1);  // Incrémenter 'page'
    }
  };

  useEffect(() => {
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);  // On attache le listener une seule fois au début

  if (loading && page === 1) {
    return <p>Chargement des produits...</p>;
  }

  return (
    <div>
      <h1>Liste des produits</h1>
      <ul>
        {products.map((product) => (
          <li key={product.id}>
            <h2>{product.name}</h2>
            <p>{product.description}</p>
            <p>Prix : ${parseFloat(product.price).toFixed(2)}</p>
            <img src={product.image_url} alt={product.name} width="100" />
          </li>
        ))}
      </ul>

      {loading && <p>Chargement...</p>}
    </div>
  );
}

export default App;
