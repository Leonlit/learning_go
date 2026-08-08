import { ReactNode } from "react";

type ListingWidgetLayoutProps<T> = {
    children: React.ReactNode;
    onNewClick: () => void;
    headers: string[]; 
    data: T[];
    newClickLabel: string;
    onSearch: (filtered: T[]) => void;
    currentPage: number;
    total: number;
    onPageChange: (page: number) => void;
};

const ListingWidgetLayout = <T,>({
    children,
    onNewClick,
    headers,
    data,
    newClickLabel,
    onSearch,
    currentPage,
    total,
    onPageChange,
}: ListingWidgetLayoutProps<T>): JSX.Element => {
    const totalPages = Math.ceil(total / 10);

    const handleSearch = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = e.target.value.toLowerCase();

        const filtered = data.filter((item: any) =>
            JSON.stringify(item).toLowerCase().includes(value),
        );

        onSearch(filtered);
    };

    const getPageNumbers = (): (number | string)[] => {
        const pages: (number | string)[] = [];

        if (totalPages <= 5) {
            return Array.from({ length: totalPages }, (_, i) => i + 1);
        }

        // Always include first page
        pages.push(1);

        if (currentPage > 3) {
            pages.push("...");
        }

        const start = Math.max(2, currentPage - 1);
        const end = Math.min(totalPages - 1, currentPage + 1);

        for (let i = start; i <= end; i++) {
            pages.push(i);
        }

        if (currentPage < totalPages - 2) {
            pages.push("...");
        }

        // Always include last page
        pages.push(totalPages);

        return pages;
    };

    const pages = getPageNumbers();
    return (
        <div className="listing-container">
            <div className="listing-menu-container">
                <button type="button" className="menu-btn" onClick={onNewClick}>
                    {newClickLabel}
                </button>

                <input
                    className="menu-input-search"
                    placeholder="Search"
                    onChange={handleSearch}
                />
            </div>

            <table className="styled-table">
                <thead>
                    <tr>
                        {headers.map((header) => (
                            <th key={header}>{header}</th>
                        ))}
                    </tr>
                </thead>
                <tbody>
                    {children}
                </tbody>
            </table>

            {totalPages > 0 && (
                <div className="pagination">
                    {currentPage > 1 && (
                        <button type="button" onClick={() => onPageChange(currentPage - 1)}>
                            Prev
                        </button>
                    )}

                    {pages.map((p, index) =>
                        p === "..." ? (
                            <span key={index} className="ellipsis">
                                ...
                            </span>
                        ) : (
                            <button
                                type="button"
                                key={index}
                                onClick={() => onPageChange(p as number)}
                                className={p === currentPage ? "active" : ""}
                            >
                                {p}
                            </button>
                        ),
                    )}

                    {currentPage < totalPages && (
                        <button 
                            type="button"
                            onClick={() => onPageChange(currentPage + 1)}
                        >
                            Next
                        </button>
                    )}
                </div>
            )}
        </div>
    );
};
export default ListingWidgetLayout;
