FROM bookshop:prod as dev.dockerfile

ENTRYPOINT task refresh --watch