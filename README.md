Test code to demonstrate https://github.com/golang/go/issues/16196 changes
==========================================================================

compressbench encodes paletted PNG images with image/png master from 2016-09-27 (8f9e2ab) and a patched copy of the png package (./png in this repo).
Writes file sizes to old.txt, new.txt and imagemagick.txt (file size from Image Magick). .txt files are benchcmp compatible (MB/s )

All test images are from golang.org/x/image/testdata/ and were converted to paletted PNG images with Image Magick:

    for i in $GOPATH/src/golang.org/x/image/testdata/*; do convert $i png8:testdata/${i##*/}.png; done


Improvements:

    file                                                                         old           new           speedup
    ./testdata/blue-purple-pink-large.lossless.webp.png                          64136         49311         0.77x
    ./testdata/blue-purple-pink-large.no-filter.lossy.webp.png                   58609         46484         0.79x
    ./testdata/blue-purple-pink-large.no-filter.lossy.webp.ycbcr.png.png         126092        121210        0.96x
    ./testdata/blue-purple-pink-large.normal-filter.lossy.webp.png               60013         46523         0.78x
    ./testdata/blue-purple-pink-large.normal-filter.lossy.webp.ycbcr.png.png     149197        144709        0.97x
    ./testdata/blue-purple-pink-large.png.png                                    64136         49311         0.77x
    ./testdata/blue-purple-pink-large.simple-filter.lossy.webp.png               58251         45929         0.79x
    ./testdata/blue-purple-pink-large.simple-filter.lossy.webp.ycbcr.png.png     135952        135233        0.99x
    ./testdata/blue-purple-pink.lossless.webp.png                                5916          4918          0.83x
    ./testdata/blue-purple-pink.lossy.webp.png                                   5377          4747          0.88x
    ./testdata/blue-purple-pink.lossy.webp.ycbcr.png.png                         12390         12424         1.00x
    ./testdata/blue-purple-pink.lzwcompressed.tiff.png                           5916          4918          0.83x
    ./testdata/blue-purple-pink.png.png                                          5916          4918          0.83x
    ./testdata/bw-deflate.tiff.png                                               539           449           0.83x
    ./testdata/bw-packbits.tiff.png                                              539           449           0.83x
    ./testdata/bw-uncompressed.tiff.png                                          539           449           0.83x
    ./testdata/go-turns-two-14x18.png.png                                        853           1108          1.30x
    ./testdata/go-turns-two-280x360.jpeg.png                                     43623         34008         0.78x
    ./testdata/go-turns-two-down-ab.png.png                                      5847          5015          0.86x
    ./testdata/go-turns-two-down-bl.png.png                                      4824          4099          0.85x
    ./testdata/go-turns-two-down-cr.png.png                                      5016          4288          0.85x
    ./testdata/go-turns-two-down-nn.png.png                                      6075          5180          0.85x
    ./testdata/go-turns-two-rotate-ab.png.png                                    2481          2130          0.86x
    ./testdata/go-turns-two-rotate-bl.png.png                                    2481          2130          0.86x
    ./testdata/go-turns-two-rotate-cr.png.png                                    2565          2276          0.89x
    ./testdata/go-turns-two-rotate-nn.png.png                                    2593          2872          1.11x
    ./testdata/go-turns-two-up-ab.png.png                                        2916          2742          0.94x
    ./testdata/go-turns-two-up-bl.png.png                                        2914          2744          0.94x
    ./testdata/go-turns-two-up-cr.png.png                                        3165          3003          0.95x
    ./testdata/go-turns-two-up-nn.png.png                                        945           1320          1.40x
    ./testdata/gopher-doc.1bpp.lossless.webp.png                                 749           579           0.77x
    ./testdata/gopher-doc.1bpp.png.png                                           749           579           0.77x
    ./testdata/gopher-doc.2bpp.lossless.webp.png                                 1158          935           0.81x
    ./testdata/gopher-doc.2bpp.png.png                                           1158          935           0.81x
    ./testdata/gopher-doc.4bpp.lossless.webp.png                                 1989          1608          0.81x
    ./testdata/gopher-doc.4bpp.png.png                                           1989          1608          0.81x
    ./testdata/gopher-doc.8bpp.lossless.webp.png                                 4700          4358          0.93x
    ./testdata/gopher-doc.8bpp.png.png                                           4700          4358          0.93x
    ./testdata/no_compress.tiff.png                                              507           526           1.04x
    ./testdata/no_rps.tiff.png                                                   507           526           1.04x
    ./testdata/testpattern.png.png                                               1016          2317          2.28x
    ./testdata/tux-rotate-ab.png.png                                             1485          1222          0.82x
    ./testdata/tux-rotate-bl.png.png                                             1779          1531          0.86x
    ./testdata/tux-rotate-cr.png.png                                             1766          1542          0.87x
    ./testdata/tux-rotate-nn.png.png                                             1391          1149          0.83x
    ./testdata/tux.lossless.webp.png                                             14718         11294         0.77x
    ./testdata/tux.png.png                                                       14718         11294         0.77x
    ./testdata/video-001-16bit.tiff.png                                          6418          5567          0.87x
    ./testdata/video-001-gray-16bit.tiff.png                                     14443         14073         0.97x
    ./testdata/video-001-gray.tiff.png                                           14443         14073         0.97x
    ./testdata/video-001-paletted.tiff.png                                       11255         10991         0.98x
    ./testdata/video-001-strip-64.tiff.png                                       6418          5567          0.87x
    ./testdata/video-001-tile-64x64.tiff.png                                     6418          5567          0.87x
    ./testdata/video-001-uncompressed.tiff.png                                   6418          5567          0.87x
    ./testdata/video-001.bmp.png                                                 6418          5567          0.87x
    ./testdata/video-001.lossy.webp.png                                          5818          5067          0.87x
    ./testdata/video-001.lossy.webp.ycbcr.png.png                                13749         13469         0.98x
    ./testdata/video-001.png.png                                                 6418          5567          0.87x
    ./testdata/video-001.tiff.png                                                6418          5567          0.87x
    ./testdata/yellow_rose-small.bmp.png                                         470           507           1.08x
    ./testdata/yellow_rose-small.png.png                                         470           507           1.08x
    ./testdata/yellow_rose.lossless.webp.png                                     22773         17192         0.75x
    ./testdata/yellow_rose.lossy-with-alpha.webp.nycbcra.png.png                 82347         70882         0.86x
    ./testdata/yellow_rose.lossy-with-alpha.webp.png                             17506         13827         0.79x
    ./testdata/yellow_rose.lossy.webp.png                                        21416         16516         0.77x
    ./testdata/yellow_rose.lossy.webp.ycbcr.png.png                              71849         62866         0.87x
    ./testdata/yellow_rose.png.png                                               22773         17192         0.75x
