/**pagination helper function */
export const paginator = (items, currentPage, perPageItems) => {
    let page = currentPage || 1,
        perPage = perPageItems || 15,
        offset = (page - 1) * perPage,
        paginatedItems = items.slice(offset).slice(0, perPageItems),
        total = items.length,
        totalPages = Math.ceil(items.length / perPage),
        from = total === 0 ? total : perPageItems * (currentPage - 1) + 1,
        to =
            perPageItems * currentPage > items.length
                ? items.length
                : perPageItems * currentPage;

    return {
        page: page,
        perPage: perPage,
        prePage: page - 1 ? page - 1 : null,
        nextPage: totalPages > page ? page + 1 : null,
        total: total,
        totalPages: totalPages,
        list: paginatedItems,
        from: from, //firstIndexOnPage
        to: to, //lastIndexOnPage
    };
};

//for processListing page
export const initialPaginationStateProcessList = {
  page: 1,
  perPageItems: 15, //fixed
  prevPage: null,
  nextPage: null,
  total: 0,
  list: [],
  from: 0,
  to: 0,
};


//for recipeListing page
export const initialPaginationStateRecipeList = {
    page: 1,
    perPageItems: 8, //fixed
    prevPage: null,
    nextPage: null,
    total: 0,
    list: [],
    from: 0,
    to: 0,
  };
  