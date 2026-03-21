import { ReactNode } from "react";

type ListingWidgetLayoutProps<T> = {
    children: React.ReactNode;
    onNewClick: () => void;
    data: T[]; // JSON array (projects, assessments, etc.)
    onSearch: (filtered: T[]) => void;
};

const ListingWidgetLayout = <T,>({
    children,
    onNewClick,
    data,
    onSearch,
}: ListingWidgetLayoutProps<T>): JSX.Element => {
    const handleSearch = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = e.target.value.toLowerCase();

        const filtered = data.filter((item: any) =>
            JSON.stringify(item).toLowerCase().includes(value),
        );

        onSearch(filtered);
    };

    return (
        <div className="listing-container">
            <div className="listing-menu-container">
                <div className="menu-btn" onClick={onNewClick}>
                    New
                </div>

                <input
                    className="menu-input-search"
                    placeholder="Search"
                    onChange={handleSearch}
                />
            </div>

            <div className="menu-item-listing">{children}</div>
        </div>
    );
};
export default ListingWidgetLayout;
