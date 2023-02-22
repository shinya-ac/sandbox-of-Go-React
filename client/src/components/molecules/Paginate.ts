import styled from '@emotion/styled';

export const PaginationContainer = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 2rem;

  .pagination {
    display: flex;
    justify-content: center;
    align-items: center;
    list-style: none;

    li {
      display: inline-block;
      margin: 0 5px;
      padding: 5px 10px;
      border: 1px solid #ccc;
      border-radius: 5px;
      cursor: pointer;

      &.active {
        background-color: #ccc;
        color: #fff;
      }

      &:hover:not(.active) {
        background-color: #f2f2f2;
      }
    }
  }
`;
