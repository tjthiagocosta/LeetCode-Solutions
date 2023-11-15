class Solution {
public:
    bool isValid(string str) {
        stack <char> s;
    
    for (int i = 0; i < str.length(); i++)
    {
        if (str[i] == '(' || str[i] == '{' || str[i] == '[')
        {
            s.push(str[i]);
        }
        else if (str[i] == ')' || str[i] == '}' || str[i] == ']')
        {
            if (s.empty() == 1) 
            {
                return false;
            }
            else if (str[i] == ')') 
            {
                char x = s.top();
                s.pop();
                if (x != '(')
                {
                    return false;
                }
            }
            else if (str[i] == '}') 
            {
                char x = s.top();
                s.pop();
                if (x != '{')
                {
                    return false;
                }
            }
            else if (str[i] == ']') 
            {
                char x = s.top();
                s.pop();
                if (x != '[')
                {
                    return false;
                }
            }
        }
    }
    if (s.empty() == 1)
    {
        return true;
    }
    else
    {
        return false;
    }
    }
};