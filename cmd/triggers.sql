CREATE OR REPLACE FUNCTION update_price_on_order()
RETURNS TRIGGER AS $$
DECLARE
    new_stock INT;
    price_change DECIMAL(5, 2);
    update_count INT DEFAULT 0;
BEGIN
    -- updating the stock level
    UPDATE products
    SET stock = stock - NEW.quantity
    WHERE product_id = NEW.product_id;

    -- checking if the updates were successful or not
    GET DIAGNOSTICS update_count = ROW_COUNT;
    IF update_count = 0 THEN
        RAISE EXCEPTION 'Product update failed for order item.';
    END IF;

    -- calculating the new stock level after update
    SELECT stock INTO new_stock
    FROM products
    WHERE product_id = NEW.product_id;

    IF new_stock < 10 THEN
        -- adjusting price if stock goes below 10
        price_change = 0.1; -- increases price by 10%
    ELSE
        IF NEW.quantity > 5 THEN
            -- adjusting price if order qty is high
            price_change = -0.05; -- decreasing the price by 5% for more sales
        ELSE
            price_change = 0;
        END IF;
    END IF;

    -- updating product price if there's a change
    IF price_change <> 0 THEN
        UPDATE products
        SET price = price * (1 + price_change)
        WHERE product_id = NEW.product_id;
    END IF;

    RETURN NULL; -- return something for triggers with AFTER type
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_price_on_order
BEFORE INSERT ON order_items
FOR EACH ROW
EXECUTE FUNCTION update_price_on_order();

CREATE OR REPLACE FUNCTION update_stock_on_order()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE products
    SET stock = stock - NEW.quantity
    WHERE product_id = NEW.product_id;

    RETURN NULL; -- return something for triggers with AFTER type
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_stock_on_order
BEFORE INSERT ON order_items
FOR EACH ROW
EXECUTE FUNCTION update_stock_on_order();