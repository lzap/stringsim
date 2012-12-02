String similarity search for Go
===============================

This simple library implements pair distance searching algorighm also known as 
2-gram method. The code was highly inspired by Ruby library amatch and 
algorithm described at http://www.catalysoft.com/articles/StrikeAMatch.html

Example use
-----------

    import "stringsim/adjpair"
    similarity := MatchStrings("test abc", "abc")
    // float value 1.0 is for same strings
    // and 0.0 is for totally different strings

There are several other functions for comparing sentences or directory paths 
which are being tokenized. Also low-level functions are exported so it is 
possible to pre-calculate pair arrays and to work faster.

License
-------

GNU LGPL v3 or later: http://www.gnu.org/copyleft/lesser.html

vim: tw=79:fo+=w
