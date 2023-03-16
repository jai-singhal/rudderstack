import React from 'react';
import { Pagination as BootstrapPagination } from 'react-bootstrap';

const Pagination = ({ currentPage, totalPages, onChangePage }) => {
	const pageItems = [];
	if(totalPages <= 1){
		return <></>
	}
  	for (let pageNumber = 1; pageNumber <= totalPages; pageNumber++) {
      pageItems.push(
		<BootstrapPagination.Item key={pageNumber} active={pageNumber===currentPage} onClick={()=> onChangePage(pageNumber)}>
			{pageNumber}
		</BootstrapPagination.Item>
		)	
	}

  	return (
		<BootstrapPagination>
		{pageItems}
		</BootstrapPagination>
  	);
};

export default Pagination;
